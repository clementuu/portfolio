package store

import "back/model"

var (
	softinnov = model.Experience{Intitule: "Développeur Fullstack", Type: "Alternance", Structure: "Softinnov", Periode: "mars 2024 à mars 2026", Taches: []string{"Conception et développement d'applications et de solutions web pour divers clients", "Administration de systèmes, migration de base de données"}}
	antiGaspi = model.Experience{Intitule: "Etude du gaspillage alimentaire", Type: "Service Civique", Structure: "Mairie de Saint-Jean de Védas", Periode: "novembre 2022 à juillet 2023", Taches: []string{"Diagnostique du gaspillage alimentaire à travers la collecte et l'étude de données", "Rédaction d'un rapport et présentation des résultats aux élus locaux", "Création d'un jeu vidéo éducatif en javascript pour sensibiliser aux écogestes"}}
	idemia    = model.Experience{Intitule: "Support sur un projet d'innovation", Type: "Stage", Structure: "IDEMIA - R&D Sophia-Antipolis", Periode: "mai 2022 à août 2022", Taches: []string{"Création et automatisation de tests d'UI", "Utilisation d'un framework de test « End to End » (Python, Selenium, Serenity, Behave)"}}
)

var experiencesList = []model.Experience{softinnov, antiGaspi, idemia}

func (r *RAMStore) GetExperiences() []model.Experience {
	return r.Experiences
}
