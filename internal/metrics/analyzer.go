package metrics

// Gerado por IA
// CDD métrics

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"path/filepath"
	"strings"
)

type FileMetrics struct {
	Filepath     string
	Couplings    int // imports de módulos do projeto
	Branches     int // if, else, switch, loops
	FunctionArgs int // funções passadas como argumento
	Total        int
	IsStateful   bool // tem struct com estado
	MaxAllowed   int
	Status       string // OK, WARNING, ERROR
}

type ProjectMetrics struct {
	Files       []FileMetrics
	ProjectPath string
	ModuleName  string
}

func (m FileMetrics) String() string {
	status := "✅ OK"
	if m.Total > m.MaxAllowed {
		if m.Total > m.MaxAllowed+3 {
			status = "🔴 COMPLEXO"
		} else {
			status = "🟡 WARNING"
		}
	}
	return fmt.Sprintf("%-40s | %d points | %s", m.Filepath, m.Total, status)
}

// AnalyzeFile analisa um arquivo .go e retorna suas métricas
func AnalyzeFile(filepath, projectName string) (*FileMetrics, error) {
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, filepath, nil, 0)
	if err != nil {
		return nil, err
	}

	m := &FileMetrics{
		Filepath:   filepath,
		IsStateful: detectStateful(file),
	}

	if m.IsStateful {
		m.MaxAllowed = 13
	} else {
		m.MaxAllowed = 7
	}

	// Contar acoplamentos (imports do projeto)
	m.Couplings = countProjectCouplings(file, projectName)

	// Contar branches e funções como argumento
	visitor := &codeVisitor{
		projectName: projectName,
	}
	ast.Walk(visitor, file)
	m.Branches = visitor.branchCount
	m.FunctionArgs = visitor.functionArgCount

	m.Total = m.Couplings + m.Branches + m.FunctionArgs

	if m.Total <= m.MaxAllowed {
		m.Status = "OK"
	} else if m.Total <= m.MaxAllowed+3 {
		m.Status = "WARNING"
	} else {
		m.Status = "ERROR"
	}

	return m, nil
}

// detectStateful detecta se arquivo tem struct com campos (state)
func detectStateful(file *ast.File) bool {
	for _, decl := range file.Decls {
		if genDecl, ok := decl.(*ast.GenDecl); ok && genDecl.Tok == token.TYPE {
			for _, spec := range genDecl.Specs {
				if typeSpec, ok := spec.(*ast.TypeSpec); ok {
					if structType, ok := typeSpec.Type.(*ast.StructType); ok {
						// Struct com campos = stateful
						if len(structType.Fields.List) > 0 {
							return true
						}
					}
				}
			}
		}
	}
	return false
}

// countProjectCouplings conta imports de módulos do projeto
func countProjectCouplings(file *ast.File, projectName string) int {
	count := 0
	for _, imp := range file.Imports {
		if imp.Path == nil {
			continue
		}
		importPath := strings.Trim(imp.Path.Value, "\"")
		// Se o import contém o nome do projeto, é um acoplamento
		if strings.Contains(importPath, projectName) {
			count++
		}
	}
	return count
}

// codeVisitor traversa a AST e conta branches e função args
type codeVisitor struct {
	branchCount      int
	functionArgCount int
	projectName      string
}

func (v *codeVisitor) Visit(node ast.Node) ast.Visitor {
	if node == nil {
		return v
	}

	// Contar if statements
	if _, ok := node.(*ast.IfStmt); ok {
		v.branchCount++
	}

	// Contar switch statements
	if _, ok := node.(*ast.SwitchStmt); ok {
		v.branchCount++
	}

	// Contar case clauses
	if _, ok := node.(*ast.CaseClause); ok {
		v.branchCount++
	}

	// Contar for/range loops
	if _, ok := node.(*ast.ForStmt); ok {
		v.branchCount++
	}
	if _, ok := node.(*ast.RangeStmt); ok {
		v.branchCount++
	}

	// Contar funções passadas como argumento
	if call, ok := node.(*ast.CallExpr); ok {
		v.countFunctionArgs(call)
	}

	return v
}

func (v *codeVisitor) countFunctionArgs(call *ast.CallExpr) {
	for _, arg := range call.Args {
		// Se argumento é uma função (FuncLit ou identificador que chama função)
		if _, ok := arg.(*ast.FuncLit); ok {
			v.functionArgCount++
		}
		// Se é uma chamada que retorna função (factory pattern)
		if callExpr, ok := arg.(*ast.CallExpr); ok {
			// Verificar se é factory pattern (func que retorna func)
			if fn, ok := callExpr.Fun.(*ast.SelectorExpr); ok {
				if _, ok := fn.X.(*ast.Ident); ok {
					// Heurística: se termina com "Handler" ou "Creator", pode ser factory
					if strings.Contains(fn.Sel.Name, "Handler") || strings.Contains(fn.Sel.Name, "Creator") {
						v.functionArgCount++
					}
				}
			}
		}
	}
}

// AnalyzeDirectory analisa todos os .go files em um diretório
func AnalyzeDirectory(dirPath, projectName string) (*ProjectMetrics, error) {
	pm := &ProjectMetrics{
		ProjectPath: dirPath,
		ModuleName:  projectName,
		Files:       []FileMetrics{},
	}

	// Encontrar todos os .go files
	files, err := filepath.Glob(filepath.Join(dirPath, "**/*.go"))
	if err != nil {
		// Fallback para diretórios simples
		files = []string{}
	}

	// Se não achou, tenta estrutura padrão
	if len(files) == 0 {
		patterns := []string{
			filepath.Join(dirPath, "cmd/*/main.go"),
			filepath.Join(dirPath, "internal/*/**.go"),
			filepath.Join(dirPath, "pkg/*/**.go"),
		}
		for _, pattern := range patterns {
			matches, _ := filepath.Glob(pattern)
			files = append(files, matches...)
		}
	}

	// Analisar cada arquivo
	for _, file := range files {
		if !strings.HasSuffix(file, ".go") || strings.Contains(file, "_test.go") {
			continue
		}

		metrics, err := AnalyzeFile(file, projectName)
		if err != nil {
			continue
		}

		pm.Files = append(pm.Files, *metrics)
	}

	return pm, nil
}

// FormatMarkdown gera relatório em markdown
func (pm *ProjectMetrics) FormatMarkdown() string {
	var sb strings.Builder

	sb.WriteString("# Análise de Complexidade — Métricas de Código\n\n")
	sb.WriteString("## Critérios\n\n")
	sb.WriteString("- **Acoplamento:** 1 ponto por import de módulo do projeto\n")
	sb.WriteString("- **Branches:** 1 ponto por if/else/switch/case/loop\n")
	sb.WriteString("- **Higher-Order Functions:** 1 ponto por função como argumento\n\n")
	sb.WriteString("**Limite:** 7 pontos (stateless) | 13 pontos (stateful)\n\n")

	sb.WriteString("## Resultados por Arquivo\n\n")
	sb.WriteString("| Arquivo | Acoplamento | Branches | F.Args | Total | Status | Max |\n")
	sb.WriteString("|---------|:-----------:|:--------:|:------:|:-----:|:------:|:---:|\n")

	for _, fm := range pm.Files {
		status := "✅"
		if fm.Status == "WARNING" {
			status = "⚠️"
		} else if fm.Status == "ERROR" {
			status = "❌"
		}

		filename := filepath.Base(fm.Filepath)
		sb.WriteString(fmt.Sprintf("| `%s` | %d | %d | %d | **%d** | %s | %d |\n",
			filename,
			fm.Couplings,
			fm.Branches,
			fm.FunctionArgs,
			fm.Total,
			status,
			fm.MaxAllowed,
		))
	}

	sb.WriteString("\n## Resumo\n\n")

	totalFiles := len(pm.Files)
	okFiles := 0
	warningFiles := 0
	errorFiles := 0

	for _, fm := range pm.Files {
		switch fm.Status {
		case "OK":
			okFiles++
		case "WARNING":
			warningFiles++
		case "ERROR":
			errorFiles++
		}
	}

	sb.WriteString(fmt.Sprintf("- **Total de arquivos:** %d\n", totalFiles))
	sb.WriteString(fmt.Sprintf("- **✅ OK:** %d\n", okFiles))
	sb.WriteString(fmt.Sprintf("- **⚠️ WARNING:** %d\n", warningFiles))
	sb.WriteString(fmt.Sprintf("- **❌ ERROR:** %d\n\n", errorFiles))

	// Total de pontos
	totalPoints := 0
	totalCouplings := 0
	totalBranches := 0
	totalFunctionArgs := 0

	for _, fm := range pm.Files {
		totalPoints += fm.Total
		totalCouplings += fm.Couplings
		totalBranches += fm.Branches
		totalFunctionArgs += fm.FunctionArgs
	}

	sb.WriteString("### Totais Projeto\n\n")
	sb.WriteString(fmt.Sprintf("| Métrica | Valor |\n"))
	sb.WriteString(fmt.Sprintf("|---------|-------|\n"))
	sb.WriteString(fmt.Sprintf("| Total Acoplamentos | %d |\n", totalCouplings))
	sb.WriteString(fmt.Sprintf("| Total Branches | %d |\n", totalBranches))
	sb.WriteString(fmt.Sprintf("| Total F.Args | %d |\n", totalFunctionArgs))
	sb.WriteString(fmt.Sprintf("| **Total Pontos** | **%d** |\n\n", totalPoints))

	// Recomendações
	if errorFiles > 0 {
		sb.WriteString("## ⚠️ Recomendações\n\n")
		sb.WriteString("Arquivos com ❌ ERROR precisam refatoração:\n")
		for _, fm := range pm.Files {
			if fm.Status == "ERROR" {
				excess := fm.Total - fm.MaxAllowed
				sb.WriteString(fmt.Sprintf("- **%s:** %d pontos acima do limite (remover %d pontos)\n",
					filepath.Base(fm.Filepath), fm.Total, excess))
			}
		}
	}

	return sb.String()
}
