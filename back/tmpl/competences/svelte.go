package tmpl

// SvelteComp contient la description HTML de la compétence Svelte.
const SvelteComp = `
<!-- =================================================================== -->
<!-- SECTION : TITRE ET INTRODUCTION DE LA COMPÉTENCE                    -->
<!-- =================================================================== -->
<section class="comp-section">
    <p>
        Svelte est un framework JavaScript moderne qui déplace le travail de l'exécution (runtime) 
        vers la compilation. Contrairement aux frameworks traditionnels comme React ou Vue, Svelte 
        compile le code en JavaScript impératif et optimisé au moment du build, ce qui permet 
        d'obtenir des applications extrêmement rapides avec une taille de bundle réduite.
    </p>
    <p>
        J'ai choisi Svelte pour le développement de l'interface de ce portfolio en raison de sa 
        simplicité, de ses performances et de son approche élégante de la réactivité. Il permet de 
        créer des composants web avec une syntaxe claire et concise, tout en offrant des 
        fonctionnalités puissantes comme le data-binding bidirectionnel et les animations intégrées.
    </p>
</section>

<!-- =================================================================== -->
<!-- SECTION : MISE EN PRATIQUE                                          -->
<!-- =================================================================== -->
<section class="comp-section">
    <h3>Mise en pratique</h3>
    <p>
        Cette compétence a été mise en œuvre sur le projet suivant et durant mon parcours professionnel.
    </p>
    <div class="project-list">
        <strong>Projets Associés</strong>
        <ul>
			<li><a class="project-link" href="/projets/detail.html?id=1">Escarcelle</a></li>       
        </ul>
    </div>
</section>
`
