package lib

import (
	"fmt"
	"reflect"
	"strings"
)

// func getGormColumnName(tag string) string {
// 	re := regexp.MustCompile(`column:([a-zA-Z0-9_]+)`)
// 	matches := re.FindStringSubmatch(tag)
// 	if len(matches) > 1 {
// 		return matches[1]
// 	}
// 	return ""
// }

// func getGormForeignKey(t reflect.Type, tag string) string {
// 	if t.Kind() != reflect.Struct {
// 		return ""
// 	}
// 	re := regexp.MustCompile(`foreignKey:([a-zA-Z0-9_]+)`)
// 	matches := re.FindStringSubmatch(tag)
// 	if len(matches) > 1 {
// 		return matches[1]
// 	}
// 	return t.Name() + "ID"
// }

// func mapTypeToPrisma(t reflect.Type) string {
// 	switch t.Kind() {
// 	case reflect.String:
// 		return "String"
// 	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
// 		return "Int"
// 	case reflect.Bool:
// 		return "Boolean"
// 	case reflect.Struct:
// 		if t == reflect.TypeOf(time.Time{}) {
// 			return "DateTime"
// 		} else {
// 			return t.Name()
// 		}
// 	default:
// 		return "String"
// 	}
// }

// func generatePrismaSchema(model interface{}, modelName string) string {
// 	schema := []string{fmt.Sprintf("model %s {", modelName)}
// 	t := reflect.TypeOf(model)

// 	for i := 0; i < t.NumField(); i++ {
// 		field := t.Field(i)
// 		prismaType := mapTypeToPrisma(field.Type)

// 		options := []string{}

// 		columnName := field.Name
// 		if gormTag, ok := field.Tag.Lookup("gorm"); ok {
// 			if strings.Contains(gormTag, "primaryKey") {
// 				options = append(options, "@id")
// 			}

// 			if strings.Contains(gormTag, "unique") {
// 				options = append(options, "@unique")
// 			}

// 			if strings.Contains(gormTag, "autoUpdateTime") {
// 				options = append(options, "@updatedAt")
// 			}

// 			fk := getGormForeignKey(t, gormTag)
// 			if fk != "" {
// 				options = append(options,
// 					fmt.Sprintf(` @relation(fields: [ %s ], references: [ id ])`, fk),
// 				)
// 			} else {
// 				customName := getGormColumnName(gormTag)
// 				if customName != "" {
// 					columnName = customName
// 				}
// 				options = append(options, fmt.Sprintf(`@map("%s")`, columnName))
// 			}
// 		}

// 		optionsString := strings.Join(options, " ")

// 		fr := []rune(field.Name)
// 		fr[0] = []rune(strings.ToLower(string(fr[0])))[0]
// 		fn := string(fr)

// 		schema = append(schema, fmt.Sprintf(`  %s %s %s`, fn, prismaType, optionsString))
// 	}

// 	_, exists := t.MethodByName("TableName")
// 	if exists {
// 		schema = append(schema, "")
// 		schema = append(schema,
// 			fmt.Sprintf(`  @@map("%s")`,
// 				model.(interface{ TableName() string }).TableName(),
// 			),
// 		)
// 	}

// 	schema = append(schema, "}")
// 	return strings.Join(schema, "\n")
// }

// func main() {
// 	u := User{}
// 	prismaSchema := generatePrismaSchema(u, "User")
// 	fmt.Println(prismaSchema)

// 	fmt.Println("")

// 	ud := UserDocument{}
// 	prismaSchema = generatePrismaSchema(ud, "UserDocument")
// 	fmt.Println(prismaSchema)
// }

// // Function to generate the Prisma schema from the given models
// func generatePrismaSchema(models ...interface{}) string {
// 	var schemaBuilder strings.Builder

// 	for _, model := range models {
// 		modelType := reflect.TypeOf(model)

// 		// Handle pointer types
// 		if modelType.Kind() == reflect.Ptr {
// 			modelType = modelType.Elem()
// 		}

// 		if modelType.Kind() != reflect.Struct {
// 			continue
// 		}

// 		// Write the model definition
// 		schemaBuilder.WriteString("model ")
// 		schemaBuilder.WriteString(modelType.Name())
// 		schemaBuilder.WriteString(" {\n")

// 		for i := 0; i < modelType.NumField(); i++ {
// 			field := modelType.Field(i)

// 			// Extract field name and type
// 			fieldName := field.Name
// 			fieldType := getPrismaType(field.Type)

// 			// Parse GORM tags for additional attributes
// 			gormTag := field.Tag.Get("gorm")
// 			fieldAttributes := parseGormTagToPrisma(gormTag)

// 			// Write field definition
// 			schemaBuilder.WriteString(fmt.Sprintf("  %s %s %s\n", fieldName, fieldType, fieldAttributes))
// 		}

// 		schemaBuilder.WriteString("}\n\n")
// 	}

// 	return schemaBuilder.String()
// }

// // Helper function to convert Go types to Prisma types
// func getPrismaType(goType reflect.Type) string {
// 	switch goType.Kind() {
// 	case reflect.String:
// 		return "String"
// 	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
// 		return "Int"
// 	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
// 		return "Int"
// 	case reflect.Float32, reflect.Float64:
// 		return "Float"
// 	case reflect.Bool:
// 		return "Boolean"
// 	case reflect.Slice:
// 		if goType.Elem().Kind() == reflect.Uint8 {
// 			return "Bytes" // For byte slices
// 		}
// 		return "String" // Simplified for array/slice types
// 	default:
// 		return "String" // Default case, should be improved for more complex types
// 	}
// }

// // Helper function to parse GORM tags and convert to Prisma attributes
// func parseGormTagToPrisma(gormTag string) string {
// 	if gormTag == "" {
// 		return ""
// 	}

// 	attributes := []string{}

// 	tagParts := strings.Split(gormTag, ";")
// 	for _, part := range tagParts {
// 		switch part {
// 		case "primaryKey":
// 			attributes = append(attributes, "@id")
// 		case "uniqueIndex":
// 			attributes = append(attributes, "@unique")
// 		case "not null":
// 			attributes = append(attributes, "@default(0)")
// 			// Add more cases as needed to parse other GORM attributes
// 		}
// 	}

// 	return strings.Join(attributes, " ")
// }

type PrismaSchemaGenerator struct {
	schemaBuilder strings.Builder
}

func (p *PrismaSchemaGenerator) Generate(models ...interface{}) string {
	for _, model := range models {
		modelType := reflect.TypeOf(model)

		// Handle pointer types
		if modelType.Kind() == reflect.Ptr {
			modelType = modelType.Elem()
		}

		if modelType.Kind() != reflect.Struct {
			continue
		}

		// Generate the model definition
		p.schemaBuilder.WriteString("model ")
		p.schemaBuilder.WriteString(modelType.Name())
		p.schemaBuilder.WriteString(" {\n")

		for i := 0; i < modelType.NumField(); i++ {
			field := modelType.Field(i)
			fieldName := field.Name
			fieldType := getPrismaType(field.Type)
			fieldAttributes := parseGormTagToPrisma(field.Tag.Get("gorm"))

			// Write the field line
			p.schemaBuilder.WriteString(fmt.Sprintf("  %s %s %s\n", fieldName, fieldType, fieldAttributes))
		}

		// Add @@map if necessary for custom table name
		tableName := getTableName(modelType)
		if tableName != modelType.Name() {
			p.schemaBuilder.WriteString(fmt.Sprintf("  @@map(\"%s\")\n", tableName))
		}

		p.schemaBuilder.WriteString("}\n\n")
	}
	return p.schemaBuilder.String()
}

// Helper function to convert Go types to Prisma types
func getPrismaType(goType reflect.Type) string {
	switch goType.Kind() {
	case reflect.String:
		return "String"
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return "Int"
	case reflect.Float32, reflect.Float64:
		return "Float"
	case reflect.Bool:
		return "Boolean"
	case reflect.Slice:
		return "String" // Simplified for slices
	default:
		return "String" // Default case
	}
}

// Helper function to parse GORM tags into Prisma attributes
func parseGormTagToPrisma(gormTag string) string {
	if gormTag == "" {
		return ""
	}

	attributes := []string{}
	tagParts := strings.Split(gormTag, ";")

	for _, part := range tagParts {
		switch part {
		case "primaryKey":
			attributes = append(attributes, "@id")
		case "uniqueIndex":
			attributes = append(attributes, "@unique")
		case "not null":
			attributes = append(attributes, "@default(0)")
		case "updatedAt":
			attributes = append(attributes, "@updatedAt")
		case "autoIncrement":
			attributes = append(attributes, "@default(autoincrement())")
		default:
			if strings.HasPrefix(part, "size:") {
				size := strings.TrimPrefix(part, "size:")
				attributes = append(attributes, fmt.Sprintf("@map(\"size:%s\")", size))
			} else if strings.HasPrefix(part, "many2many:") {
				joinTable := strings.TrimPrefix(part, "many2many:")
				attributes = append(attributes, fmt.Sprintf("@relation(fields: [id], references: [id]) @map(\"%s\")", joinTable))
			}
		}
	}

	return strings.Join(attributes, " ")
}

// Helper function to get custom table name from GORM tags
func getTableName(t reflect.Type) string {
	gormTableName := t.Name() // Default to struct name
	// Custom logic for table name (if needed, could parse struct tags or metadata here)
	return gormTableName
}
