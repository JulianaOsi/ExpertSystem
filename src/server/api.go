package server

import (
	"bytes"
	"context"
	"encoding/json"
	"expertSystem/src/store"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"html/template"
	"io/ioutil"
	"log"
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
		Diagnosis *store.Diagnosis
	}
	var v viewData
	var err error
	vars := mux.Vars(r)
	knowledgeId, err := strconv.Atoi(vars["knowledge"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logrus.Errorf("failed to convert string to int: %v\n", err)
		return
	}

	symptom, err := strconv.Atoi(vars["symptom"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logrus.Errorf("failed to convert string to int: %v\n", err)
		return
	}

	v.Knowledge, err = store.DB.GetKnowledgeByInnerId(context.Background(), symptom, knowledgeId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logrus.Errorf("failed to get root knowledge: %v\n", err)
		return
	}

	if v.Knowledge.DiagnosisId != nil {
		v.Diagnosis, err = store.DB.GetDiagnosisById(context.Background(), *v.Knowledge.DiagnosisId)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			logrus.Errorf("failed to get diagnosis by id: %v\n", err)
			return
		}

		temp, err := template.ParseFiles("src/server/templates/conclusion.html")
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
		Amount    int
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

	k, err := store.DB.GetAllKnowledge(context.Background())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logrus.Errorf("failed to get all knowledge: %v\n", err)
		return
	}

	v.Amount = len(k)

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
	type knowledge struct {
		Ks []store.Knowledge `json:"knowledges"`
	}
	vars := mux.Vars(r)
	symptomId, _ := strconv.Atoi(vars["symptom"])
	k := knowledge{}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalf("failed to read body: %v\n", err)
	}

	err = json.Unmarshal(body, &k)
	if err != nil {
		log.Fatalf("failed to unmarshal json: %v\n", err)
	}

	for i, _ := range k.Ks {
		k.Ks[i].SymptomId = symptomId
		err := store.DB.CreateKnowledge(context.Background(), k.Ks[i])
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			logrus.Errorf("failed to create knowledge: %v\n", err)
			return
		}
	}

	r, err = http.NewRequest(http.MethodGet, "/knowledge", bytes.NewReader([]byte("")))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logrus.Errorf("failed to create new request: %v\n", err)
		return
	}
	http.Redirect(w, r, "/knowledge", http.StatusSeeOther)
}
