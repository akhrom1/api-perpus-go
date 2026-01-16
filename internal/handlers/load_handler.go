package handlers

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"api-perpus-go/internal/services"

	"github.com/gin-gonic/gin"
)

type loanReq struct {
	MemberID int   `form:"member_id" json:"member_id"`
	BookIDs  string `form:"book_ids" json:"book_ids"`
	DueDate  string `form:"due_date" json:"due_date"` 
}

func CreateLoan(c *gin.Context) {
	var req loanReq
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	var bookIDs []int
	ids := strings.Split(req.BookIDs, ",")
	for _, id := range ids {
		id = strings.TrimSpace(id) 
		num, err := strconv.Atoi(id)
		if err != nil {
			c.JSON(400, gin.H{"message": "invalid book_ids"})
			return
		}
		bookIDs = append(bookIDs, num)
	}

	var due time.Time


	if req.DueDate != "" {
		var err error
		due, err = time.Parse("2006-01-02", req.DueDate)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "invalid due_date format (use YYYY-MM-DD)"})
			return
		}
	}


	data, err := services.CreateLoan(req.MemberID, bookIDs, due)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, data)
}


func ReturnLoan(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid id"})
		return
	}

	if err := services.ReturnLoan(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "returned"})
}


func GetLoans(c *gin.Context) {
	data, err := services.GetAllLoans()
	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}
	c.JSON(200, data)
}

func GetLoanDetail(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := services.GetLoanByID(id)
	if err != nil {
		c.JSON(404, gin.H{"message": "loan not found"})
		return
	}
	c.JSON(200, data)
}
