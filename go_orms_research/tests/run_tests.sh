#!/bin/sh

echo "🚀 Go ORM Testing Environment Setup"
echo "==================================="

# Colors for output
GREEN='\033[0;32m'
BLUE='\033[0;34m'
YELLOW='\033[1;33m'
RED='\033[0;31m'
NC='\033[0m' # No Color

# Check if we're in the right directory
if [ ! -f "go.mod" ]; then
    echo -e "${RED}❌ Error: go.mod not found. Please run this script from the code/tests directory.${NC}"
    exit 1
fi

echo -e "${BLUE}📦 Installing dependencies...${NC}"
go mod tidy

echo -e "${BLUE}🧪 Testing individual ORM samples...${NC}"

# Vet and test all samples
echo -e "${YELLOW}Running go vet and go test...${NC}"
(cd samples/gorm && go vet ./... && go test ./...)
(cd samples/ent && go vet ./... && go test ./...)
(cd samples/sqlc && go vet ./... && go test ./...)
echo ""

# Test GORM sample
echo -e "${YELLOW}Testing GORM sample:${NC}"
(cd samples/gorm && CGO_ENABLED=1 go run main.go)
echo ""

# Test Ent sample
echo -e "${YELLOW}Testing Ent sample:${NC}"
(cd samples/ent && CGO_ENABLED=1 go run main.go)
echo ""

# Test SQLC sample
echo -e "${YELLOW}Testing SQLC sample:${NC}"
(cd samples/sqlc && CGO_ENABLED=1 go run main.go)
echo ""

echo -e "${BLUE}⚡ Running performance benchmarks...${NC}"
go run test_runner.go

echo -e "${GREEN}✅ All tests completed successfully!${NC}"
echo ""
echo -e "${BLUE}📁 Code samples are available in:${NC}"
echo "  - code/samples/gorm/"
echo "  - code/samples/ent/"
echo "  - code/samples/sqlc/"
echo ""
echo -e "${BLUE}🔧 To regenerate code:${NC}"
echo "  - For Ent: cd ../samples/ent && go generate ./..."
echo "  - For SQLC: cd ../samples/sqlc && sqlc generate" 