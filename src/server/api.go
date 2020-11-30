package server

import (
	"bytes"
	"context"
	"expertSystem/src/store"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"html/template"
	"net/http"
	"strconv"
)

func welcomeHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "src/server/templates/welcome.html")
}

func symptomHandler(w http.ResponseWriter, r *http.Request) {
	type viewData struct {
		Symptoms []*store.Symptom
	}
	var v viewData
	var err error
	v.Symptoms, err = store.DB.GetAllSymptoms(context.Background())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logrus.Errorf("failed to get all symptoms: %v\n", err)
		return
	}

	temp, err := template.ParseFiles("src/server/templates/symptoms.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logrus.Errorf("failed to parse html: %v\n", err)
		return
	}
	buf := new(bytes.Buffer)
	err = temp.Execute(buf, v)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logrus.Errorf("failed to execute template: %v\n", err)
		return
	}

	_, err = w.Write(buf.Bytes())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logrus.Errorf("failed to wrute response: %v\n", err)
		return
	}
}

func testStartHandler(w http.ResponseWriter, r *http.Request) {
	type viewData struct {
		Knowledge *store.Knowledge
		Question  *store.Question
	}
	var v viewData
	var err error
	vars := mux.Vars(r)
	symptomId, err := strconv.Atoi(vars["symptom"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logrus.Errorf("failed to convert string to int: %v\n", err)
		return
	}

	v.Knowledge, err = store.DB.GetRootKnowledge(context.Background(), symptomId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logrus.Errorf("failed to get root knowledge: %v\n", err)
		return
	}

	v.Question, err = store.DB.GetQuestionsById(context.Background(), v.Knowledge.QuestionId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logrus.Errorf("failed to get question by id: %v\n", err)
		return
	}

	temp, err := template.ParseFiles("src/server/templates/test.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logrus.Errorf("failed to parse html: %v\n", err)
		return
	}
	buf := new(bytes.Buffer)
	err = temp.Execute(buf, v)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logrus.Errorf("failed to execute template: %v\n", err)
		return
	}

	_, err = w.Write(buf.Bytes())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logrus.Errorf("failed to wrute response: %v\n", err)
		return
	}

}

func testHandler(w http.ResponseWriter, r *http.Request) {
	type viewData struct {
		Knowledge *store.Knowledge
		Question  *store.Question
	}
	var v viewData
	var err error
	vars := mux.Vars(r)
	symptomId, err := strconv.Atoi(vars["knowledge"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logrus.Errorf("failed to convert string to int: %v\n", err)
		return
	}

	v.Knowledge, err = store.DB.GetKnowledgeById(context.Background(), symptomId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logrus.Errorf("failed to get root knowledge: %v\n", err)
		return
	}

	v.Question, err = store.DB.GetQuestionsById(context.Background(), v.Knowledge.QuestionId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logrus.Errorf("failed to get question by id: %v\n", err)
		return
	}

	temp, err := template.ParseFiles("src/server/templates/test.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logrus.Errorf("failed to parse html: %v\n", err)
		return
	}
	buf := new(bytes.Buffer)
	err = temp.Execute(buf, v)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logrus.Errorf("failed to execute template: %v\n", err)
		return
	}

	_, err = w.Write(buf.Bytes())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logrus.Errorf("failed to write response: %v\n", err)
		return
	}
}

func knowledgeSectionHandler(w http.ResponseWriter, r *http.Request) {
	type viewData struct {
		Symptoms []*store.Symptom
	}
	var v viewData
	var err error
	v.Symptoms, err = store.DB.GetAllSymptoms(context.Background())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logrus.Errorf("failed to get all symptoms: %v\n", err)
		return
	}

	temp, err := template.ParseFiles("src/server/templates/knowledgeSection.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logrus.Errorf("failed to parse html: %v\n", err)
		return
	}
	buf := new(bytes.Buffer)
	err = temp.Execute(buf, v)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logrus.Errorf("failed to execute template: %v\n", err)
		return
	}

	_, err = w.Write(buf.Bytes())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logrus.Errorf("failed to wrute response: %v\n", err)
		return
	}
}

func knowledgeHandler(w http.ResponseWriter, r *http.Request) {
	type viewData struct {
		Questions []*store.Question
		Knowledge []*store.Knowledge
		Diagnosis []*store.Diagnosis
	}
	var v viewData
	var err error
	vars := mux.Vars(r)
	symptomId, err := strconv.Atoi(vars["symptom"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logrus.Errorf("failed to convert string to int: %v\n", err)
		return
	}

	v.Questions, err = store.DB.GetQuestionsBySymptomId(context.Background(), symptomId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logrus.Errorf("failed to get question by symptom id: %v\n", err)
		return
	}

	v.Knowledge, err = store.DB.GetKnowledgeBySymptomId(context.Background(), symptomId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logrus.Errorf("failed to get knowledge by symptom id: %v\n", err)
		return
	}

	v.Diagnosis, err = store.DB.GetDiagnosisBySymptomId(context.Background(), symptomId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logrus.Errorf("failed to get diagnosis by symptom id: %v\n", err)
		return
	}

	temp, err := template.ParseFiles("src/server/templates/knowledge.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logrus.Errorf("failed to parse html: %v\n", err)
		return
	}
	buf := new(bytes.Buffer)
	err = temp.Execute(buf, v)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logrus.Errorf("failed to execute template: %v\n", err)
		return
	}

	_, err = w.Write(buf.Bytes())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logrus.Errorf("failed to write response: %v\n", err)
		return
	}
}

func addKnowledgeHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	symptomId, _ := strconv.Atoi(vars["symptom"])

	var knowledge store.Knowledge
	isRoot, _ := strconv.ParseBool(r.FormValue("is_root"))
	question, _ := strconv.Atoi(r.FormValue("question"))
	trueQ, _ := strconv.Atoi(r.FormValue("true"))
	falseQ, _ := strconv.Atoi(r.FormValue("false"))
	diagnosis, _ := strconv.Atoi(r.FormValue("diagnosis"))

	knowledge.SymptomId = symptomId
	knowledge.IsRoot = isRoot
	knowledge.QuestionId = &question
	knowledge.TrueQuestionId = &trueQ
	knowledge.FalseQuestionId = &falseQ
	knowledge.DiagnosisId = &diagnosis

	err := store.DB.CreateKnowledge(context.Background(), knowledge)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logrus.Errorf("failed to create knowledge: %v\n", err)
		return
	}

	r, err = http.NewRequest(http.MethodGet, "/knowledge", bytes.NewReader([]byte("")))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logrus.Errorf("failed to create new request: %v\n", err)
		return
	}
	http.Redirect(w, r, "/knowledge", http.StatusSeeOther)
}
