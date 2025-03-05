package validation

import (
	. "cv-file-validate/models"
	"errors"
	"fmt"
)

func ValidateResume(resume Resume) error {
	var validationErrors []string

	if resume.Name == "" {
		validationErrors = append(validationErrors, "missing required field: name")
	}
	if resume.Initials == "" {
		validationErrors = append(validationErrors, "missing required field: initials")
	}
	if resume.Summary == "" {
		validationErrors = append(validationErrors, "missing required field: summary")
	}
	if resume.About == "" {
		validationErrors = append(validationErrors, "missing required field: about")
	}
	if resume.Contact.Email == "" {
		validationErrors = append(validationErrors, "missing required field: contact.email")
	}
	if resume.Contact.Tel == "" {
		validationErrors = append(validationErrors, "missing required field: contact.tel")
	}
	if len(resume.Education) == 0 {
		validationErrors = append(validationErrors, "missing required field: education")
	}
	for i, edu := range resume.Education {
		if edu.School == "" {
			validationErrors = append(validationErrors, fmt.Sprintf("missing required field: education[%d].school", i))
		}
		if edu.Degree == "" {
			validationErrors = append(validationErrors, fmt.Sprintf("missing required field: education[%d].degree", i))
		}
		if edu.Start == "" {
			validationErrors = append(validationErrors, fmt.Sprintf("missing required field: education[%d].start", i))
		}
		if edu.End == "" {
			validationErrors = append(validationErrors, fmt.Sprintf("missing required field: education[%d].end", i))
		}
	}
	if len(resume.Work) == 0 {
		validationErrors = append(validationErrors, "missing required field: work")
	}
	for i, job := range resume.Work {
		if job.Company == "" {
			validationErrors = append(validationErrors, fmt.Sprintf("missing required field: work[%d].company", i))
		}
		if job.Title == "" {
			validationErrors = append(validationErrors, fmt.Sprintf("missing required field: work[%d].title", i))
		}
		if job.Start == "" {
			validationErrors = append(validationErrors, fmt.Sprintf("missing required field: work[%d].start", i))
		}
		if job.End == "" {
			validationErrors = append(validationErrors, fmt.Sprintf("missing required field: work[%d].end", i))
		}
		if job.Description == "" {
			validationErrors = append(validationErrors, fmt.Sprintf("missing required field: work[%d].description", i))
		}
	}
	if len(resume.Skills) == 0 {
		validationErrors = append(validationErrors, "missing required field: skills")
	}
	for i, project := range resume.Projects {
		if project.Title == "" {
			validationErrors = append(validationErrors, fmt.Sprintf("missing required field: projects[%d].title", i))
		}
		if len(project.TechStack) == 0 {
			validationErrors = append(validationErrors, fmt.Sprintf("missing required field: projects[%d].techStack", i))
		}
		if project.Description == "" {
			validationErrors = append(validationErrors, fmt.Sprintf("missing required field: projects[%d].description", i))
		}
	}

	if len(validationErrors) > 0 {
		return errors.New("validation errors:\n" + fmt.Sprintf("%v", validationErrors))
	}

	return nil
}

func IsValidYAML(filename string) bool {
	return len(filename) > 4 && (filename[len(filename)-5:] == ".yaml" || filename[len(filename)-4:] == ".yml")
}
