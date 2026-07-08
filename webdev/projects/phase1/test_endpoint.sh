
# Exercise 1: /method-inspector
echo -e "${BLUE}[Exercise 1: /method-inspector]${NC}"
R1=$(curl -s -X GET "$SERVER_URL/method-inspector")
if [[ "$R1" == *"GET"* ]]; then echo -e "${GREEN}✔ PASS: GET detected${NC}"; else echo -e "${RED}✘ FAIL: got '$R1'${NC}"; fi
R1P=$(curl -s -X POST "$SERVER_URL/method-inspector")
if [[ "$R1P" == *"POST"* ]]; then echo -e "${GREEN}✔ PASS: POST detected${NC}"; else echo -e "${RED}✘ FAIL: got '$R1P'${NC}"; fi
