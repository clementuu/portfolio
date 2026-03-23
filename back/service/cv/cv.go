package cv

import "back/model"

type CV struct {
	Formations  []model.Formation
	Experiences []model.Experience
}

func GetCV() CV {
	return CV{
		Formations:  storage.GetFormations(),
		Experiences: storage.GetExperiences(),
	}
}
