
// Example curl command line to define index in the CouchDB channel_chaincode database
// curl -i -X POST -H "Content-Type: application/json" -d "{"index":{"fields":["objectType","creationDate"]},"ddoc":"indexcreationDateDoc", "name":"indexcreationDate","type":"json"}" http://hostname:port/myc1_marbles/_index


// peer chaincode invoke -C $CHANNEL_NAME -n maincc -c '{"Args":["createDocument", "Orascomproject", "testRFRstring", "Peer1", "Peer2", "TestBody", "testAttachment", "TestattachType"]}'
// peer chaincode invoke -C $CHANNEL_NAME -n maincc -c '{"Args":["readDocument", "Orascomproject" ]}'
// peer chaincode invoke -C $CHANNEL_NAME -n maincc -c '{"Args":["deleteDocument", "IBMproject" ]}'

// peer chaincode invoke -C $CHANNEL_NAME -n maincc -c '{"Args":["queryDocumentByDate", CreationDate ]}'
// peer chaincode invoke -C $CHANNEL_NAME -n maincc -c '{"Args":["getDocumentHistory", Subject ]}'

// peer chaincode invoke -C $CHANNEL_NAME -n maincc -c '{"Args":["modifiyDocument","Dellproject", "newSSSSSSSSSSSSSSSSAttachment", "AttachmentTyesayweywtwype", "newSexyBody" ]}'
// peer chaincode invoke -C $CHANNEL_NAME -n maincc -c '{"Args":["modifiyOfferState", Subject, "Rejected" ]}'

// peer chaincode invoke -C $CHANNEL_NAME -n maincc -c '{"Args":["addComment", Subject, commenter, comment ]}'
// peer chaincode invoke -C $CHANNEL_NAME -n maincc -c '{"Args":["getAllDocuments"]}'


// Rich Query with index design doc and index name specified (Only supported if CouchDB is used as state database):
//   peer chaincode query -C myc1 -n marbles -c '{"Args":["queryMarbles","{\"selector\":{\"objectType\":\"marble\",\"owner\":\"tom\"}, \"use_index\":[\"_design/indexOwnerDoc\", \"indexOwner\"]}"]}'

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/hyperledger/fabric-chaincode-go/shim"
	pb "github.com/hyperledger/fabric-protos-go/peer"
)



type Chaincode struct {

}

type Comment struct {
	Comment string 	`json:"comment"`
	User string		`json:"user"`
} 

type Document struct {
	ObjectType string 		`json:"objectType"`
	Subject string			`json:"subject"`
	DocumentType	string	`json:"documentType"`

	Sender string			`json:"sender"`
	Receiver string 		`json:"receiver"`

	CreationDate string 	`json:"creationDate"`
	CreationTime string 	`json:"creationTime"`

	Body               string 	`json:"body"`
	Attachment         string 	`json:"attachment"`
	AttachmentType     string	`json:"attachmentType"`
	 

	// Confirmation state is for offers only (Confirmed, Rejected, Waiting for confirmation  )
	ConfirmationState  string 	`json:"confirmationState"`


	Modified           bool   	`json:"modified"`
	LastModified 	   string	`json:"lastModified"`

	Comments 		 []Comment	`json:"comments"`

}


// ===================================================================================
// Main
// ===================================================================================

func main() {
    err := shim.Start(new(Chaincode))
	if err != nil {
		fmt.Printf("Error starting  Chaincode: %s", err)
	}
}


// Init initializes chaincode
// ===========================

func (t *Chaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	return shim.Success(nil)
}

// Invoke - Invoke method for chaincode
// ========================================

func (t *Chaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
    
    function, args := stub.GetFunctionAndParameters()

	fmt.Println("invoke is running " + function)

	// Handle different functions
    if function == "createDocument" {
		return t.createDocument(stub, args)
			
    } else if function == "readDocument" {
		return t.readDocument(stub, args)
		
	} else if function == "deleteDocument" {
		return t.deleteDocument(stub, args)
			
	}  else if function == "queryDocumentByDate" {
		return t.queryDocumentByDate(stub, args)

	} else if function == "getDocumentHistory" {
		return t.getDocumentHistory(stub,args)

	} else if function == "modifiyDocument" {
		return t.modifiyDocument(stub, args)

	} else if function  == "modifiyOfferState" {
		return t.modifiyOfferState(stub, args)
 
	} else if function  == "addComment" {
		return t.addComment(stub, args)

	} else if function  == "getAllDocuments" {
		return t.getAllDocuments(stub)

	}

    fmt.Println("invoke did not find func: " + function) //error
	return shim.Error("Received unknown function invocation")
}


// ============================================================
// createDocument - create a new document, store into chaincode state
// ============================================================

func (t *Chaincode) createDocument(stub shim.ChaincodeStubInterface, args []string) pb.Response {
		var err error
		 
		if len(args) != 7 {
			return shim.Error("Incorrect number of arguments. Expecting 7")
		}
	
		// ==== Input sanitation ====
		if len(args[0]) <= 0 {
			return shim.Error("Subject argument must be a non-empty string")
		}
		if len(args[1]) <= 0 {
			return shim.Error("DocumentType argument must be a non-empty string")
		}
		if len(args[2]) <= 0 {
			return shim.Error("Sender argument must be a non-empty string")
		}
		if len(args[3]) <= 0 {
			return shim.Error("Receiver argument must be a non-empty string")
		}
		if len(args[4]) <= 0 {
			return shim.Error("Body argument must be a non-empty string")
		}
		if len(args[5]) <= 0 {
			return shim.Error("Attachment argument must be a non-empty string")
		}
		if len(args[6]) <= 0 {
			return shim.Error("Attachment Type argument must be a non-empty string")
		}
		
		
	
		fmt.Println(" ******** Creating Document  ******** ")
	
		subject := args [0]
		documentType := args[1]

		sender := args[2]
		receiver := args[3]

		body := args[4]
		attachment := args[5]
		attachmentType := args[6]

		
		

		// ==== Check if this document already exists ====
		documentAsBytes, err := stub.GetState(subject)
		if err != nil {
			return shim.Error("Failed to get document: " + err.Error())
		} else if documentAsBytes != nil {
			fmt.Println("This document already exists: " + subject)
			return shim.Error("This document already exists: " + subject)
		}
	
		// ==== Create document object and marshal to JSON ====
		objectType := "document"
		currentTime := time.Now()
		creationDate := currentTime.Format("01-02-2006")
		creationTime := currentTime.Format("3:4:5 PM")
		var comments []Comment
		

		document := &Document{objectType, subject, documentType, sender,
		receiver, creationDate, creationTime,
		body, attachment, attachmentType, "", false, "", comments}

		documentJSONasBytes, err := json.Marshal(document)

		if err != nil {
		return shim.Error(err.Error())
	}		
	
			err = stub.PutState(subject, documentJSONasBytes)
			if err !=nil {
				return shim.Error(err.Error())
			}

	
	indexName := "subject~creationDate"
	dateSubjectIndexKey, err := stub.CreateCompositeKey(indexName, []string{document.Subject, document.CreationDate})
	if err != nil {
		return shim.Error(err.Error())
	}
	//  Save index entry to state. Only the key name is needed, no need to store a duplicate copy of the Document.
	//  Note - passing a 'nil' value will effectively delete the key from state, therefore we pass null character as value
	value := []byte{0x00}
	stub.PutState(dateSubjectIndexKey, value)

	// ==== document saved and indexed. Return success ====
	
	fmt.Println("- end creating document")
	return shim.Success(nil)
}

// ===============================================
// readDocument - read a document from chaincode state
// ===============================================

func (t *Chaincode) readDocument(stub shim.ChaincodeStubInterface, args []string) pb.Response {
		var subject, jsonResp string
		var err error

		if len(args) != 1 {
			return shim.Error("Incorrect number of arguments. Expecting subject of the document to query")
		}

		subject = args[0]
		valAsbytes, err := stub.GetState(subject) //get the document from chaincode state

		if err != nil {
			jsonResp = "{\"Error\":\"Failed to get state for " + subject + "\"}"
			return shim.Error(jsonResp)
		} else if valAsbytes == nil {
			jsonResp = "{\"Error\":\"subject does not exist: " + subject + "\"}"
			return shim.Error(jsonResp)
		}

		return shim.Success(valAsbytes)
}	


// ==================================================
// delete - remove a dockument key/value pair from state
// ==================================================

func (t *Chaincode) deleteDocument(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var jsonResp string
	var documentJSON Document
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}
	subject := args[0]

	// to maintain the subject~creationDate index, we need to read the document first and get its creationDate
	valAsbytes, err := stub.GetState(subject) //get the document from chaincode state
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + subject + "\"}"
		return shim.Error(jsonResp)
	} else if valAsbytes == nil {
		jsonResp = "{\"Error\":\"Document does not exist: " + subject + "\"}"
		return shim.Error(jsonResp)
	}

	err = json.Unmarshal([]byte(valAsbytes), &documentJSON)
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to decode JSON of: " + subject + "\"}"
		return shim.Error(jsonResp)
	}

	err = stub.DelState(subject) //remove the document from chaincode state
	if err != nil {
		return shim.Error("Failed to delete state:" + err.Error())
	}

	// maintain the index
	indexName := "subject~creationDate"
	dateSubjectIndexKey, err := stub.CreateCompositeKey(indexName, []string{documentJSON.Subject, documentJSON.CreationDate})
	if err != nil {
		return shim.Error(err.Error())
	}

	//  Delete index entry to state.
	err = stub.DelState(dateSubjectIndexKey)
	if err != nil {
		return shim.Error("Failed to delete state:" + err.Error())
	}
	return shim.Success(nil)
}

// ==================================================
// queryDocumentByDate 
// ==================================================

func (t *Chaincode) queryDocumentByDate(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) < 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	creationDate := args[1]

	queryString := fmt.Sprintf("{\"selector\":{\"objectType\":\"document\",\"creationDate\":\"%s\"}}", creationDate)

	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(queryResults) 
}

func getQueryResultForQueryString(stub shim.ChaincodeStubInterface, queryString string) ([]byte, error) {

	fmt.Printf("- getQueryResultForQueryString queryString:\n%s\n", queryString)

	resultsIterator, err := stub.GetQueryResult(queryString)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	buffer, err := constructQueryResponseFromIterator(resultsIterator)
	if err != nil {
		return nil, err
	}

	fmt.Printf("- getQueryResultForQueryString queryResult:\n%s\n", buffer.String())

	return buffer.Bytes(), nil
}

func constructQueryResponseFromIterator(resultsIterator shim.StateQueryIteratorInterface) (*bytes.Buffer, error) {
	// buffer is a JSON array containing QueryResults
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.Key)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Record\":")
		// Record is a JSON object, so we write as-is
		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	return &buffer, nil
}


func (t *Chaincode) getDocumentHistory(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) < 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	subject := args[0]

	fmt.Printf("- start getDocumentHistory: %s\n",subject)

	resultsIterator, err := stub.GetHistoryForKey(subject)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing historic values for the document
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		response, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"TxId\":")
		buffer.WriteString("\"")
		buffer.WriteString(response.TxId)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Value\":")
		// if it was a delete operation on given key, then we need to set the
		//corresponding value null. Else, we will write the response.Value
		//as-is (as the Value itself a JSON Document)
		if response.IsDelete {
			buffer.WriteString("null")
		} else {
			buffer.WriteString(string(response.Value))
		}

		buffer.WriteString(", \"Timestamp\":")
		buffer.WriteString("\"")
		buffer.WriteString(time.Unix(response.Timestamp.Seconds, int64(response.Timestamp.Nanos)).String())
		buffer.WriteString("\"")

		buffer.WriteString(", \"IsDelete\":")
		buffer.WriteString("\"")
		buffer.WriteString(strconv.FormatBool(response.IsDelete))
		buffer.WriteString("\"")

		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("- getHistoryForDocument returning:\n%s\n", buffer.String())

	return shim.Success(buffer.Bytes())
}


// ===========================================================
// Modifiy Document Attachment
// ===========================================================

func (t *Chaincode) modifiyDocument(stub shim.ChaincodeStubInterface, args []string) pb.Response {


	if len(args) < 4 {
		return shim.Error("Incorrect number of arguments. Expecting 4")
	}

	subject := args[0]
	newAttachment := args[1]
	newAttachmentType := args[2]
	newBody := args[3]
	fmt.Println("- Start modifying document ")

	documentAsBytes, err := stub.GetState(subject)
	if err != nil {
		return shim.Error("Failed to get document:" + err.Error())
	} else if documentAsBytes == nil {
		return shim.Error("document does not exist")
	}

	documentToModifiy := Document{}

	err = json.Unmarshal(documentAsBytes, &documentToModifiy) //unmarshal it aka JSON.parse()
	if err != nil {
		return shim.Error(err.Error())
	}
	documentToModifiy.Attachment = newAttachment //change the attachment 
	documentToModifiy.AttachmentType = newAttachmentType	//change the attachmentType 
	documentToModifiy.Body = newBody //change the body

	documentToModifiy.Modified = true
	currentTime := time.Now()
	documentToModifiy.LastModified = currentTime.Format("2006-01-02 3:4:5 PM")

	documentJSONasBytes, _ := json.Marshal(documentToModifiy)
	err = stub.PutState(subject, documentJSONasBytes) //rewrite the Document
	if err != nil {
		return shim.Error(err.Error())
	}

	fmt.Println("- end transferDocument (success)")
	return shim.Success(nil)
}

// ===========================================================
// Modifiy Offer State
// ===========================================================

func (t *Chaincode) modifiyOfferState(stub shim.ChaincodeStubInterface, args []string) pb.Response {


	if len(args) < 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	subject := args[0]
	confirmationState := args[1]
	
	fmt.Println("- Start modifying document ")

	documentAsBytes, err := stub.GetState(subject)

	if err != nil {
		return shim.Error("Failed to get document:" + err.Error())
	} else if documentAsBytes == nil {
		return shim.Error("document does not exist")
	}

	documentToModifiy := Document{}

	err = json.Unmarshal(documentAsBytes, &documentToModifiy) //unmarshal it aka JSON.parse()
	if err != nil {
		return shim.Error(err.Error())
	}

	if documentToModifiy.DocumentType != "Offer"{
		return shim.Error("document is not an offer")
	}

	documentToModifiy.ConfirmationState = confirmationState //change the confirmationState 
	

	documentJSONasBytes, _ := json.Marshal(documentToModifiy)
	err = stub.PutState(subject, documentJSONasBytes) //rewrite the Document
	if err != nil {
		return shim.Error(err.Error())
	}

	fmt.Println("- end transferDocument (success)")
	return shim.Success(nil)
}


// ===========================================================
// Add Comment 
// ===========================================================

func (t *Chaincode) addComment(stub shim.ChaincodeStubInterface, args []string) pb.Response {


	if len(args) < 3 {
		return shim.Error("Incorrect number of arguments. Expecting 3")
	}

	subject := args[0]
	commenter := args[1]
	comment := args[2] 

	
	fmt.Println("- Start modifying document ")

	documentAsBytes, err := stub.GetState(subject)
	if err != nil {
		return shim.Error("Failed to get document:" + err.Error())
	} else if documentAsBytes == nil {
		return shim.Error("document does not exist")
	}

	documentToModifiy := Document{}

	err = json.Unmarshal(documentAsBytes, &documentToModifiy) //unmarshal it aka JSON.parse()
	if err != nil {
		return shim.Error(err.Error())
	}
	documentToModifiy.Comments = append( documentToModifiy.Comments, Comment{Comment:comment, User:commenter}) //change the confirmationState 
	

	documentJSONasBytes, _ := json.Marshal(documentToModifiy)
	err = stub.PutState(subject, documentJSONasBytes) //rewrite the Document
	if err != nil {
		return shim.Error(err.Error())
	}

	fmt.Println("- end transferDocument (success)")
	return shim.Success(nil)
}

// ===========================================================
// Get All documents 
// ===========================================================

func (t *Chaincode) getAllDocuments(stub shim.ChaincodeStubInterface) pb.Response {
	
	startKey := ""
	endKey := ""

	resultsIterator, err := stub.GetStateByRange(startKey, endKey)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	buffer, err := constructQueryResponseFromIterator(resultsIterator)
	if err != nil {
		return shim.Error(err.Error())
	}

	fmt.Printf("- getAllDocuments queryResult:\n%s\n", buffer.String())

	return shim.Success(buffer.Bytes())
}