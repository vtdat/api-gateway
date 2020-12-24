package controllers

import (
	"benchmark/api-gateway/models"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestHandler_caclQuantile(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	r := gin.Default()
	handler := NewHandler()
	v1 := r.Group("/api/v1")
	handler.MakeHandler(v1)
	var b io.Reader = nil

	primes := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes[1:4])
	superPrimes := [1000]int{
		2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97,
		2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97,
		2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97,
		2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97,
		2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97,
		2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97,
		2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97,
		2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97,
		2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97,
		2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97,
		2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97,
		2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97,
	}
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		args args
		req  *models.Quantile
		want map[string]interface{}
	}{
		{
			name: "test invalid request - case 1",
			args: args{
				c: c,
			},
			req: &models.Quantile{
				Pool:       make([]int, 0),
				Percentile: 99.5,
			},
			want: map[string]interface{}{
				"code":        float64(400),
				"message":     nil,
				"description": "No value in the pool",
			},
		},
		{
			name: "test invalid request - case 2",
			args: args{
				c: c,
			},
			req: &models.Quantile{
				Pool:       primes,
				Percentile: 199.5,
			},
			want: map[string]interface{}{
				"code":        float64(400),
				"message":     nil,
				"description": "Percentile must be float type in range (0-100]",
			},
		},
		{
			name: "test valid request - case 5",
			args: args{
				c: c,
			},
			req: &models.Quantile{
				Pool:       primes[1:4],
				Percentile: float64(0.1),
			},
			want: map[string]interface{}{
				"code":        float64(200),
				"message":     float64(3),
				"description": "The point which 0.1% of values in the pool are less than or equal is 3",
			},
		},
		{
			name: "test valid request - case 6",
			args: args{
				c: c,
			},
			req: &models.Quantile{
				Pool:       primes[1:4],
				Percentile: 100,
			},
			want: map[string]interface{}{
				"code":        float64(200),
				"message":     float64(7),
				"description": "The point which 100% of values in the pool are less than or equal is 7",
			},
		},
		{
			name: "test valid request - case 7",
			args: args{
				c: c,
			},
			req: &models.Quantile{
				Pool:       primes[1:4],
				Percentile: 50,
			},
			want: map[string]interface{}{
				"code":        float64(200),
				"message":     float64(5),
				"description": "The point which 50% of values in the pool are less than or equal is 5",
			},
		},
		{
			name: "test valid request - case 8",
			args: args{
				c: c,
			},
			req: &models.Quantile{
				Pool:       superPrimes[0:1000],
				Percentile: 50,
			},
			want: map[string]interface{}{
				"code":        float64(400),
				"message":     nil,
				"description": "Simulate DDOS attack! Request over 1000 values is not acceptable",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if data, _ := json.Marshal(tt.req); data != nil {
				b = bytes.NewBuffer(data)
			} else {
				t.Fatalf("cannot marshall data")
			}
			w = httptest.NewRecorder()

			req, err := http.NewRequest("POST", "/api/v1/statistics/quantile", b)
			if err != nil {
				t.Fatalf("cannot POST data")
			}
			req.Header.Add("Content-Type", "application/json")
			r.ServeHTTP(w, req)

			if w.Code != int(tt.want["code"].(float64)) {
				t.Fatalf("Wrong output, got %d, want %v", w.Code, tt.want["code"])
			}


			fmt.Println(string(w.Body.Bytes()))
			var got map[string]interface{}
			err = json.Unmarshal(w.Body.Bytes(), &got)
			if err != nil {
				t.Fatal(err)
			}


			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("Wrong output, got %v, want %v", got, tt.want)
			}
			w.Flush()
			w.Body.Reset()
		})
	}
}
