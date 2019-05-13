package main

import(
  "github.com/gorilla/mux"
  "fmt"
  "log"
  "net/http"
  "strconv"
  "github.com/aws/aws-sdk-go/aws"
  "github.com/aws/aws-sdk-go/aws/session"
  "github.com/aws/aws-sdk-go/aws/awserr"
  "github.com/aws/aws-sdk-go/service/dynamodb"

)

func main(){
  r:= mux.NewRouter()
  r.HandleFunc("/dkear/test", TestFunc)
  r.HandleFunc("/dkear/all", All)
  r.HandleFunc("/dkear/status", Status )

  http.Handle("/", r)
  log.Fatal(http.ListenAndServe(":9170", nil))
}

func TestFunc(w http.ResponseWriter, r *http.Request){
  w.WriteHeader(http.StatusOK)
  fmt.Fprint(w, "testtest")
}

func All(w http.ResponseWriter, r *http.Request){
  w.WriteHeader(http.StatusOK)
  sess, err := session.NewSession(&aws.Config{Region: aws.String("us-east-1")},)
  svc := dynamodb.New(sess)
  input := &dynamodb.ScanInput{TableName: aws.String("dkearR6"),
    }

    result, err := svc.Scan(input)
    if err != nil {
      if aerr, ok := err.(awserr.Error); ok {
        switch aerr.Code(){
        case dynamodb.ErrCodeResourceNotFoundException:
          fmt.Println(dynamodb.ErrCodeResourceNotFoundException, aerr.Error())
        case dynamodb.ErrCodeInternalServerError:
          //fmt.Println(dynamodb.ErrCodeINternalServerError, aerr.Error())
        default:
          fmt.Println(aerr.Error())
        }
      }else{
          fmt.Println(err.Error())
        }
        return
      }
      fmt.Fprint(w, result)
}

func Status(w http.ResponseWriter, r *http.Request){
  w.WriteHeader(http.StatusOK)
  sess, err := session.NewSession(&aws.Config{Region: aws.String("us-east-1")},)
  svc := dynamodb.New(sess)
  input := &dynamodb.DescribeTableInput{TableName: aws.String("dkearR6"),
    }
    result, err := svc.DescribeTable(input)
    if err != nil {
    if aerr, ok := err.(awserr.Error); ok {
        switch aerr.Code() {
        case dynamodb.ErrCodeResourceNotFoundException:
            fmt.Println(dynamodb.ErrCodeResourceNotFoundException, aerr.Error())
        case dynamodb.ErrCodeInternalServerError:
            //fmt.Println(dynamodb.ErrCodeInternalServerError, aerr.Error())
        default:
            fmt.Println(aerr.Error())
        }
    } else {
        fmt.Println(err.Error())
    }
        return
    }

    data := "{\"table\": \"dkearR6\", \"ItemCount\":" + strconv.FormatInt(*(result.Table.ItemCount), 10) + "}"
    fmt.Fprint(w, data)
}
