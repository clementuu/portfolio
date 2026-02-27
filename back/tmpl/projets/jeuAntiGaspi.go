package tmpl

const JeuAntiGaspiHTML = `
<!-- =================================================================== -->
<!-- SECTION : TITRE ET INTRODUCTION DU PROJET                           -->
<!-- =================================================================== -->
<section class="project-section">
    <p>
        Ce projet est un jeu vidéo 2D conçu pendant mon service civique pour la mairie de 
        Saint-Jean-de-Védas. L'objectif était de sensibiliser les enfants aux éco-gestes et à la 
        réduction du gaspillage alimentaire à travers une expérience ludique, simple et accessible.
    </p>

    <p>
        Le jeu a été développé entièrement en <a class="link" href="/competences/detail.html?id=3">JavaScript</a> 
        et <a class="link" href="/competences/detail.html?id=2">HTML/CSS</a> sans framework ni moteur 
        externe. Il s'agit d'un jeu en vue du dessus dans lequel le joueur doit accomplir différentes 
        actions liées au tri, à la gestion des déchets ou à la consommation responsable. 
        Ce projet représente une étape importante dans mon parcours : c'est l'un de mes premiers 
        développements interactifs, et il m'a permis d'apprendre les bases de la programmation 
        orientée jeu vidéo.
    </p>
</section>

<!-- =================================================================== -->
<!-- SECTION : ACTEURS ET INTERACTIONS DU PROJET                        -->
<!-- =================================================================== -->
<section class="project-section">
    <h3><b>Acteurs et Interactions</b></h3>
    <p>
        Ce projet a été réalisé dans un contexte particulier : j'étais le seul développeur et testeur, 
        chargé de concevoir l'intégralité du jeu, de son gameplay à son intégration graphique. 
        Les interactions se faisaient principalement avec :
    </p>

    <ul>
        <li>
            <strong>Moi-même (développeur et testeur)</strong> - responsable de la conception, du développement, 
            des tests, de l'équilibrage du gameplay et de l'intégration des contenus pédagogiques.
        </li>

        <li>
            <strong>La mairie de Saint-Jean-de-Védas</strong> - commanditaire du projet, définissant les objectifs 
            pédagogiques et validant les mécaniques liées aux éco-gestes.
        </li>

        <li>
            <strong>Les enfants (public cible)</strong> - leurs retours informels ont permis d'ajuster la difficulté, 
            la lisibilité des actions et la dimension ludique du jeu.
        </li>
    </ul>

    <p>
        Cette configuration m'a permis de travailler en autonomie complète, tout en gardant un lien 
        constant avec les besoins pédagogiques du commanditaire et les attentes du jeune public.
    </p>
</section>

<!-- =================================================================== -->
<!-- SECTION : MISSIONS ET COMPÉTENCES DÉVELOPPÉES                       -->
<!-- =================================================================== -->
<section class="project-section">
    <h3><b>Missions et Compétences Développées</b></h3>
    <p>
        Ce projet m'a permis de découvrir et d'expérimenter les fondations du développement de jeux 
        2D en <a class="link" href="/competences/detail.html?id=3">JavaScript</a>. 
        J'ai dû apprendre à manipuler des ressources graphiques, gérer un canvas, 
        créer une boucle de jeu fluide et mettre en place des mécaniques interactives.
    </p>

    <div class="skills-grid">

        <div class="skill-item">
            <strong>Canvas HTML5 & Contexte 2D</strong>
            <p>
                Utilisation du canvas pour dessiner les éléments du jeu, gérer les rafraîchissements 
                d'écran, afficher le personnage, les objets interactifs et les animations simples.
            </p>
        </div>

        <div class="skill-item">
            <strong>Game Loop & Animation</strong>
            <p>
                Mise en place d'une boucle de jeu personnalisée (gameloop) pour gérer les mises à jour 
                du gameplay, les déplacements, les collisions et le rendu graphique à chaque frame.
            </p>
        </div>

        <div class="skill-item">
            <strong>Gestion des Ressources (Images & Sons)</strong>
            <p>
                Chargement et intégration d'images sprites, d'éléments de décor et d'effets sonores. 
                Gestion du préchargement et optimisation simple pour éviter les ralentissements.
            </p>
        </div>

        <div class="skill-item">
            <strong>Interactions & Collisions</strong>
            <p>
                Détection de collisions entre le joueur et les objets du décor, zones d'interaction 
                permettant de déclencher des actions (ramasser, trier, déposer…).  
                Mise en place de règles de jeu simples pour guider l'enfant dans l'apprentissage des 
                bons gestes.
            </p>
        </div>

        <div class="skill-item">
            <strong>Logique de Jeu & Pédagogie</strong>
            <p>
                Conception de mécaniques ludiques adaptées aux enfants : objectifs clairs, feedbacks 
                visuels et sonores, progression simple.  
                Intégration de messages éducatifs sur le tri, la consommation responsable et la 
                réduction du gaspillage.
            </p>
        </div>

        <div class="skill-item">
            <strong>Autonomie & Apprentissage</strong>
            <p>
                Projet réalisé sans moteur de jeu ni librairie externe, nécessitant une compréhension 
                approfondie du fonctionnement interne du canvas, des événements clavier et de la 
                gestion du temps.  
                Ce projet a été un excellent terrain d'apprentissage et reste une réalisation dont je 
                suis fier.
            </p>
        </div>

    </div>
</section>

<!-- =================================================================== -->
<!-- SECTION : CONCLUSION                                                -->
<!-- =================================================================== -->
<section class="project-section">
    <h3><b>Conclusion</b></h3>
    <p>
        Ce jeu antigaspi est un projet qui me tient particulièrement à cœur. Bien que son code ne 
        reflète plus mon niveau actuel, il représente une étape fondatrice dans mon parcours de 
        développeur : j'y ai appris les bases du développement de jeux 2D, la manipulation du canvas, 
        la gestion d'assets et la création d'interactions ludiques.  
        Surtout, il a permis de sensibiliser des enfants aux éco-gestes de manière amusante et 
        pédagogique, ce qui en fait une réussite autant technique qu'humaine.
    </p>

    <p>
        Les lendemains du projet restent ouverts : aucune évolution fonctionnelle n'est prévue pour 
        l'instant, mais j'envisage une refonte complète afin d'améliorer la qualité du code, 
        moderniser la structure et faciliter la maintenance.  
        Je pourrais également profiter de mon serveur VPS pour rendre le jeu accessible en ligne, 
        afin qu'il puisse continuer à être utilisé dans un cadre éducatif ou associatif.
    </p>
</section>

<!-- =================================================================== -->
<!-- SECTION : COMPÉTENCES ASSOCIÉES                                     -->
<!-- =================================================================== -->
<section class="project-section">
    <h5>Compétences Associées</h5>
    <div class="competences-list">
        <a href="/competences/detail.html?id=3" class="competence-tag technique">JavaScript</a>
        <a href="/competences/detail.html?id=2" class="competence-tag technique">HTML/CSS</a>
        <a href="/competences/detail.html?id=15" class="competence-tag humain">Créativité</a>
    </div>
</section>
`
