package main

import (
	"fmt"
	"os"
)

var soutab = [...]string{"DS_COUNTER_GROUP_NK_CI_L_6M",
	"DS_COUNTER_GROUP_NK_CM_L_6M",
	"DS_COUNTER_GROUP_NK_CS_L_6M",
	"DS_COUNTER_GROUP_NK_CG_L_6M",
	"DS_COUNTER_GROUP_NK_CY_L_6",
	"DS_COUNTER_GROUP_NK_CB_L_6M",
	"DS_COUNTER_GROUP_NK_CC_L_6M",
	"DS_COUNTER_GROUP_NK_CY_L_6M",
	"DS_COUNTER_GROUP_NK_CM_L_6Q",
	"DS_COUNTER_GROUP_NK_CS_L_6Q",
	"DS_COUNTER_GROUP_NK_CY_L_6Q",
	"DS_COUNTER_GROUP_NK_CI_L_6Q",
	"DS_COUNTER_GROUP_NK_CB_L_6Q",
	"DS_COUNTER_GROUP_NK_CC_L_6Q",
	"DS_COUNTER_GROUP_NK_CG_L_6Q",
	"DS_COUNTER_GROUP_NK_CS_L_6",
	"DS_COUNTER_GROUP_NK_CC_L_6D",
	"DS_COUNTER_GROUP_NK_CG_L_6D",
	"DS_COUNTER_GROUP_NK_CI_L_6D",
	"DS_COUNTER_GROUP_NK_CB_L_6D",
	"DS_COUNTER_GROUP_NK_CB_L_6B",
	"DS_COUNTER_GROUP_NK_CC_L_6B",
	"DS_COUNTER_GROUP_NK_CM_L_6B",
	"DS_COUNTER_GROUP_NK_CM_L_6D",
	"DS_COUNTER_GROUP_NK_CG_L_6",
	"DS_COUNTER_GROUP_NK_CI_L_6",
	"DS_COUNTER_GROUP_NK_CM_L_6",
	"DS_COUNTER_GROUP_NK_CC_L_6",
	"DS_COUNTER_GROUP_NK_CS_L_6D",
	"DS_COUNTER_GROUP_NK_CY_L_6D",
	"DS_COUNTER_GROUP_NK_CB_L_6",
}

var col = [...]string{"M8201C20",
	"M8205C0",
	"M8205C1",
	"M8205C2",
	"M8205C3",
	"M8205C4",
	"M8205C5",
	"M8205C6",
	"M8205C7",
	"M8205C8",
	"M8205C9",
	"M8205C10",
	"M8205C11",
	"M8205C12",
	"M8205C13",
	"M8205C14",
	"M8205C15",
	"M8205C16",
	"M8205C17",
	"M8205C18",
	"M8205C19",
	"M8205C20",
	"M8205C21",
	"M8205C22",
	"M8205C23",
	"M8205C24",
	"M8205C25",
	"M8205C26",
	"M8205C27",
	"M8205C28",
	"M8205C29",
	"M8205C30",
	"M8205C31",
	"M8205C32",
	"M8204C0",
	"M8204C1",
	"M8204C8",
	"M8204C9",
	"M8204C22",
	"M8204C77",
	"M8204C78",
	"M8204C79",
	"M8204C80",
	"M8204C88",
	"M8204C89",
	"M8204C90",
	"M8204C91",
	"M8204C92",
	"M8204C93",
	"M8204C94",
	"M8204C95",
	"M8204C96",
	"M8204C97",
	"M8204C98",
	"M8204C56",
	"M8204C57",
	"M8204C58",
	"M8204C74",
	"M8204C75",
	"M8204C76",
	"M8204C2",
	"M8204C3",
	"M8204C4",
	"M8204C5",
	"M8204C6",
	"M8204C7",
	"M8204C10",
	"M8204C11",
	"M8204C12",
	"M8204C13",
	"M8204C14",
	"M8204C15",
	"M8204C16",
	"M8204C17",
	"M8204C18",
	"M8204C19",
	"M8204C20",
	"M8204C21",
	"M8204C23",
	"M8204C24",
	"M8204C25",
	"M8204C32",
	"M8204C33",
	"M8204C34",
	"M8204C35",
	"M8204C36",
	"M8204C37",
	"M8204C38",
	"M8204C39",
	"M8204C40",
	"M8204C41",
	"M8204C42",
	"M8204C43",
	"M8204C44",
	"M8204C45",
	"M8204C46",
	"M8204C47",
	"M8204C48",
	"M8204C49",
	"M8204C50",
	"M8204C51",
	"M8204C52",
	"M8204C53",
	"M8204C54",
	"M8204C55",
	"M8204C59",
	"M8204C60",
	"M8204C61",
	"M8204C62",
	"M8204C63",
	"M8204C64",
	"M8204C65",
	"M8204C66",
	"M8204C67",
	"M8204C68",
	"M8204C69",
	"M8204C81",
	"M8204C82",
	"M8204C83",
	"M8204C84",
	"M8204C85",
	"M8204C86",
	"M8204C87",
	"M8204C99",
	"M8204C100",
	"M8204C101",
	"M8204C102",
	"M8204C103",
	"M8204C104",
	"M8204C105",
	"M8204C106",
	"M8204C107",
	"M8204C108",
	"M8204C109",
	"M8204C110",
	"M8204C111",
	"M8204C112",
	"M8204C113",
	"M8204C114",
	"M8204C115",
	"M8204C116",
	"M8204C117",
	"M8204C118",
	"M8204C119",
	"M8204C120",
	"M8204C121",
	"M8204C122",
	"M8200C33",
	"M8200C34",
	"M8200C35",
	"M8200C36",
	"M8200C0",
	"M8200C1",
	"M8200C2",
	"M8200C3",
	"M8200C4",
	"M8200C5",
	"M8200C6",
	"M8200C7",
	"M8200C8",
	"M8200C9",
	"M8200C10",
	"M8200C11",
	"M8200C12",
	"M8200C13",
	"M8200C14",
	"M8200C15",
	"M8200C16",
	"M8200C17",
	"M8200C18",
	"M8200C19",
	"M8200C20",
	"M8200C21",
	"M8200C22",
	"M8200C23",
	"M8200C24",
	"M8200C25",
	"M8200C26",
	"M8200C27",
	"M8200C28",
	"M8200C29",
	"M8200C30",
	"M8200C31",
	"M8200C32",
	"M8200C37",
	"M8200C38",
	"M8200C39",
	"M8200C40",
	"M8200C41",
	"M8200C42",
	"M8200C43",
	"M8200C44",
	"M8200C45",
	"M8200C46",
	"M8200C47",
	"M8200C48",
	"M8200C49",
	"M8200C50",
	"M8200C51",
	"M8201C0",
	"M8201C1",
	"M8201C2",
	"M8201C3",
	"M8201C4",
	"M8201C5",
	"M8201C6",
	"M8201C7",
	"M8201C8",
	"M8201C9",
	"M8201C10",
	"M8201C11",
	"M8201C12",
	"M8201C13",
	"M8201C14",
	"M8201C15",
	"M8201C16",
	"M8201C17",
	"M8201C18",
	"M8201C19",
}

func tab() {
	sqlFile := "sqlscript.sql"
	f, err := os.OpenFile(sqlFile, os.O_WRONLY|os.O_TRUNC, 0600)
	fmt.Println(err)
	defer f.Close()

	fmt.Print(len(soutab))
	for _, v := range soutab {
		fmt.Print("--", v, "\n")
		for _, x := range col {
			sql := "alter table " + v + " add " + x + " number;"
			fmt.Print(sql, "\n")
			f.WriteString(sql)
		}

	}
}
