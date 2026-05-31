package ssac

const sampleSSaC = `package course

import "context"

type GetCourseRequest struct {
	ID int64
}

// @get Course course = Course.FindByID({ID: request.id})
// @response course
func GetCourse(ctx context.Context, request GetCourseRequest) {}
`
