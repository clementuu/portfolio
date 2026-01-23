package tmpl

// WebComp contient la description HTML de la compétence Web (HTML/CSS).
const WebComp = `
<!-- =================================================================== -->
<!-- SECTION : TITRE ET INTRODUCTION DE LA COMPÉTENCE                    -->
<!-- =================================================================== -->
<section class="comp-section">
    <p>
        HTML (HyperText Markup Language) et CSS (Cascading Style Sheets) sont les piliers fondamentaux du développement web. 
        HTML fournit la structure sémantique du contenu, tandis que CSS est responsable de la présentation, du style et de la mise en page. 
        La maîtrise de ces deux langages est essentielle pour créer des sites web accessibles, réactifs et esthétiquement agréables.
    </p>
    <p>
        Mon expérience en HTML/CSS me permet de concevoir des interfaces utilisateur intuitives et modernes, 
        en veillant à ce que l'expérience soit cohérente sur tous les appareils et navigateurs. 
        Je porte une attention particulière à la sémantique HTML pour l'accessibilité (a11y) et le référencement (SEO), 
        et j'utilise des techniques CSS avancées comme Flexbox et Grid pour des mises en page complexes.
    </p>
</section>

<!-- =================================================================== -->
<!-- SECTION : MISE EN PRATIQUE                                          -->
<!-- =================================================================== -->
<section class="comp-section">
    <h3>Mise en pratique</h3>
    <p>
        Mes compétences en développement web ont été mises en œuvre sur des projets variés, démontrant ma capacité à concevoir et réaliser des expériences utilisateur riches et fonctionnelles, allant de l'application de gestion à l'outil ludo-éducatif.
    </p>
    <ul>
        <li>
            <strong>Développement d'applications web dynamiques :</strong> Conception et réalisation d'interfaces complexes et modulaires avec des frameworks modernes comme <strong>Svelte</strong> (pour l'application de gestion Escarcelle) et <strong>Angular</strong> (pour l'outil de gestion de projet PMT). Cela inclut la création d'architectures à base de composants, la gestion de l'état applicatif et la communication asynchrone avec des API REST.
        </li>
        <li>
            <strong>Développement interactif en JavaScript natif :</strong> Création d'une expérience ludo-éducative (Jeu Anti-Gaspi) entièrement en <strong>HTML, CSS et JavaScript "vanilla"</strong>. Ce projet a nécessité la manipulation directe du <strong>Canvas HTML5</strong>, la mise en place d'une boucle de jeu, la gestion des animations et la détection de collisions.
        </li>
        <li>
            <strong>Ergonomie et expérience utilisateur (UI/UX) :</strong> Un accent particulier a été mis sur la conception d'interfaces claires, ergonomiques et responsives. Que ce soit pour définir des workflows utilisateurs fluides dans <strong>Escarcelle</strong> ou pour créer une navigation intuitive dans <strong>PMT</strong>, l'objectif est toujours de fournir une expérience utilisateur de qualité.
        </li>
    </ul>
</section>
<section class="comp-section">
    <h3>Projets Associés</h3>
    <div class="project-list">
        <ul>
            <li><a class="project-link" href="/projets/detail.html?id=1">Escarcelle</a></li>
            <li><a class="project-link" href="/projets/detail.html?id=4">Jeu Anti-Gaspi</a></li>
            <li><a class="project-link" href="/projets/detail.html?id=6">PMT/a></li>
        </ul>
    </div>
</section>
`
