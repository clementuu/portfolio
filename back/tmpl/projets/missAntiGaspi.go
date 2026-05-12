package tmpl

const MissionAntiGaspiHTML = `
<!-- =================================================================== -->
<!-- SECTION : TITRE ET INTRODUCTION DU PROJET                           -->
<!-- =================================================================== -->
<section class="project-section">
    <h3><b>Présentation</b></h3>

    <p>
        <b>TousAntiGaspi</b> est une mission menée dans le cadre de mon service civique pour la mairie 
        de Saint-Jean-de-Védas. L'objectif était de réaliser une analyse fine du gaspillage alimentaire 
        dans les écoles de la commune afin d'en identifier les causes principales et de proposer des 
        leviers d'action concrets aux élus locaux.
    </p>

    <p>
        Cette étude a nécessité la mise en place d'un protocole rigoureux de collecte de données, 
        l'analyse statistique de plusieurs centaines de pesées réalisées sur trois années, et la 
        production d'un rapport complet présenté en comité de pilotage.  
        Le projet mêlait travail de terrain, rigueur scientifique et vulgarisation auprès d'un public 
        non spécialiste.
    </p>
</section>

<section class="project-section">
    <h3><b>Contexte et Enjeux</b></h3>
    <p>
        En 2021, la commune faisait face à un constat préoccupant : près de <b>46 %</b> de la nourriture 
        livrée par le traiteur n'était pas consommée, soit en moyenne <b>268 g de gaspillage par enfant 
        et par repas</b>. Les causes étaient multiples : livraisons trop importantes, écarts entre repas 
        commandés et besoins réels, manque de tri, faible sensibilisation des enfants, conditions 
        d'accueil difficiles et absence d'accompagnement au repas.
    </p>

    <p>
        La municipalité a donc lancé un projet pluriannuel visant à réduire durablement le gaspillage 
        alimentaire. Ma mission s'inscrivait dans cette démarche : analyser les données issues des 
        pesées réalisées dans les écoles, comprendre les mécanismes du gaspillage et mesurer l'impact 
        des actions mises en place, notamment le passage d'un menu à 5 composantes à un menu à 4 
        composantes.
    </p>

    <p>
        Les enjeux étaient à la fois environnementaux, économiques et éducatifs : réduire les déchets, 
        optimiser les commandes, améliorer la qualité du temps de repas et sensibiliser les enfants 
        aux enjeux alimentaires.  
        L'étude devait fournir des résultats fiables pour orienter les décisions publiques et préparer 
        le futur marché de restauration scolaire.
    </p>

    <h3><b>Acteurs et Interactions</b></h3>
    <p>
        La mission TousAntiGaspi a mobilisé deux acteurs principaux aux rôles clairement définis :
    </p>

    <ul>
        <li>
            <strong>Moi-même</strong> - responsable de la conception du protocole, 
            de la collecte des données, de l'analyse statistique, de la rédaction du rapport et 
            de la présentation des résultats aux décideurs.
        </li>

        <li>
            <strong>La mairie de Saint-Jean-de-Védas</strong> - définissant les objectifs 
            de l'étude, facilitant l'accès aux écoles, validant les orientations méthodologiques 
            et recevant les recommandations finales pour mise en œuvre.
        </li>
    </ul>

    <p>
        Les interactions ont été régulières tout au long de la mission : réunions de cadrage, 
        validation du protocole, échanges sur les premiers résultats et restitution finale auprès 
        des élus et des services scolaires. Cette collaboration étroite a permis d'aligner 
        l'analyse scientifique avec les besoins opérationnels de la collectivité.
    </p>
</section>


<!-- =================================================================== -->
<!-- SECTION : MISSIONS ET COMPÉTENCES DÉVELOPPÉES                       -->
<!-- =================================================================== -->
<section class="project-section">
    <h3><b>Missions et Compétences Développées</b></h3>
    <p>
        Cette mission m'a permis de combiner travail de terrain, analyse statistique avancée et 
        communication scientifique. J'ai mené l'ensemble du processus, de la collecte des données 
        jusqu'à la restitution finale en comité de pilotage.
    </p>

    <div class="skills-grid">

        <div class="skill-item">
            <strong>Collecte et Structuration des Données</strong>
            <p>
                Mise en place d'un protocole de pesées dans plusieurs écoles (12 à 18 pesées par session).  
                Catégorisation des composantes (entrée, plat, laitage, fruit, dessert, pain…).  
                Nettoyage, consolidation et préparation des données pour l'analyse.
            </p>
        </div>

        <div class="skill-item">
            <strong>Analyse Statistique (R & Python)</strong>
            <p>
                Étude de l'évolution du gaspillage entre 2021 et 2023 (baisse de 268 g à 177 g).  
                Tests statistiques (Wilcoxon) pour mesurer l'impact du passage à 4 composantes.  
                Analyse en Composantes Principales (ACP) pour identifier les composantes les plus 
                gaspillées (fruits et entrées en tête).
            </p>
        </div>

        <div class="skill-item">
            <strong>Visualisation et Interprétation</strong>
            <p>
                Création de graphiques, tableaux et représentations visuelles pour faciliter la 
                compréhension des résultats.  
                Interprétation des axes factoriels de l'ACP et traduction en constats opérationnels.
            </p>
        </div>

        <div class="skill-item">
            <strong>Analyse des Impacts</strong>
            <p>
                Mesure de l'impact du passage à 4 composantes :
                <ul>  
                    <li>Baisse significative du gaspillage (p-value = 0,02)</li>  
                    <li>Aucune baisse de la consommation des enfants (p-value = 0,26)</li> 
                Analyse du gaspillage par type de composante et identification des recettes problématiques.
                </ul>    
            </p>
        </div>

        <div class="skill-item">
            <strong>Rédaction et Présentation du Rapport</strong>
            <p>
                Rédaction d'un rapport complet présenté en comité de pilotage (COPIL).  
                Vulgarisation des résultats pour les élus et les services scolaires.  
                Formulation de recommandations concrètes : tri, redistribution, adaptation des menus, 
                formation des équipes, boîtes de satisfaction, etc.
            </p>
        </div>

        <div class="skill-item">
            <strong>Impact Social et Environnemental</strong>
            <p>
                Contribution directe à la réduction du gaspillage alimentaire dans les écoles.  
                Participation à la sensibilisation des enfants et des équipes éducatives.  
                Aide à la préparation du futur marché public de restauration scolaire.
            </p>
        </div>

    </div>
</section>

<!-- =================================================================== -->
<!-- SECTION : CONCLUSION                                                -->
<!-- =================================================================== -->
<section class="project-section">
    <h3><b>Résultats</b></h3>
    <p>
        L'étude a permis de mettre en évidence une baisse significative du gaspillage alimentaire 
        entre 2021 et 2023, passant de 268 g à 177 g par enfant et par repas.  
        Le passage à 4 composantes a eu un impact positif sur la réduction du gaspillage, sans 
        diminuer la consommation des enfants.  
        Les analyses ont également permis d'identifier les composantes les plus problématiques 
        (fruits et entrées) et de proposer des pistes d'action concrètes.
    </p>

    <p>
        Pour moi, cette mission a été une expérience riche mêlant statistiques, travail de terrain, 
        communication et engagement citoyen. Elle m'a permis de développer une expertise solide en 
        analyse de données appliquée à des enjeux réels.
    </p>

    <h3><b>Pour le futur</b></h3>
    <p>
        Les résultats de l'étude ont servi de base à la réflexion sur le futur marché public de 
        restauration scolaire.  
        La mairie poursuit désormais les actions engagées : tri des biodéchets, redistribution, 
        adaptation des menus, sensibilisation des enfants, amélioration des commandes et suivi 
        régulier des pesées.
    </p>

    <h3><b>Mon analyse critique</b></h3>
    <p>
        Avec le recul, cette mission a été une réussite : les données ont permis d'éclairer des 
        décisions publiques et d'améliorer concrètement les pratiques dans les écoles.  
        J'aurais toutefois pu aller plus loin dans l'automatisation des analyses ou dans la création 
        d'outils de suivi pour les équipes.  
        Malgré cela, ce projet reste l'un des plus formateurs de mon parcours, mêlant rigueur 
        scientifique et impact social tangible.
    </p>
</section>

<!-- =================================================================== -->
<!-- SECTION : COMPÉTENCES ASSOCIÉES                                     -->
<!-- =================================================================== -->
<section class="project-section">
    <h5>Compétences Associées</h5>
    <div class="competences-list">
        <a href="/competences/detail.html?id=10" class="competence-tag technique">R</a>
        <a href="/competences/detail.html?id=8" class="competence-tag technique">Python</a>
        <a href="/competences/detail.html?id=16" class="competence-tag humain">Esprit critique</a>
    </div>
</section>
`
