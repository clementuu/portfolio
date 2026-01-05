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
        Le jeu a été développé entièrement en <b>HTML / CSS / JavaScript</b> sans framework ni moteur 
        externe. Il s'agit d'un jeu en vue du dessus dans lequel le joueur doit accomplir différentes 
        actions liées au tri, à la gestion des déchets ou à la consommation responsable. 
        Ce projet représente une étape importante dans mon parcours : c'est l'un de mes premiers 
        développements interactifs, et il m'a permis d'apprendre les bases de la programmation 
        orientée jeu vidéo.
    </p>
</section>

<!-- =================================================================== -->
<!-- SECTION : MISSIONS ET COMPÉTENCES DÉVELOPPÉES                       -->
<!-- =================================================================== -->
<section class="project-section">
    <h3>Missions et Compétences Développées</h3>
    <p>
        Ce projet m'a permis de découvrir et d'expérimenter les fondations du développement de jeux 
        2D en JavaScript. J'ai dû apprendre à manipuler des ressources graphiques, gérer un canvas, 
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
    <h3>Conclusion</h3>
    <p>
        Ce jeu antigaspi est un projet qui me tient particulièrement à cœur. Bien que son code ne 
        reflète plus mon niveau actuel, il représente une étape fondatrice dans mon parcours de 
        développeur : j'y ai appris les bases du développement de jeux 2D, la manipulation du canvas, 
        la gestion d'assets et la création d'interactions ludiques.  
        Surtout, il a permis de sensibiliser des enfants aux éco-gestes de manière amusante et 
        pédagogique, ce qui en fait une réussite autant technique qu'humaine.
    </p>
</section>
<!-- =================================================================== -->
<!-- SECTION : COMPÉTENCES ASSOCIÉES                                     -->
<!-- =================================================================== -->
<section class="project-section">
    <h3>Compétences Associées</h3>
    <div class="competences-list">
        <a href="/competences/detail.html?id=3" class="competence-tag technique">JavaScript</a>
        <a href="/competences/detail.html?id=2" class="competence-tag technique">HTML/CSS</a>
        <a href="/competences/detail.html?id=17" class="competence-tag humain">Créativité</a>
    </div>
</section>
`
