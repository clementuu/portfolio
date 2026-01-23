package tmpl

// BashComp contient la description HTML de la compétence Bash.
const BashComp = `
<!-- =================================================================== -->
<!-- SECTION : TITRE ET INTRODUCTION DE LA COMPÉTENCE                    -->
<!-- =================================================================== -->
<section class="comp-section">
    <p>
        Bash (Bourne-Again SHell) est l'interpréteur de commandes par défaut sur la plupart des systèmes Linux. Au-delà d'une simple interface en ligne de commande, c'est un langage de script puissant qui permet d'automatiser une grande variété de tâches système et de développement.
    </p>
    <p>
        Utilisant Linux comme environnement de travail principal depuis plus de deux ans, j'ai rapidement adopté Bash comme un outil essentiel de mon quotidien. Je l'utilise aussi bien pour des commandes simples que pour l'écriture de scripts d'automatisation.
    </p>
</section>

<!-- =================================================================== -->
<!-- SECTION : MISE EN PRATIQUE                                          -->
<!-- =================================================================== -->
<section class="comp-section">
    <h3>Mise en pratique</h3>
    <p>
        Mon utilisation de Bash va de la simple gestion de mon environnement de développement à la création de scripts pour des opérations complexes et critiques, notamment dans le cadre du projet <strong>Escarcelle</strong>.
    </p>
    <ul>
        <li>
            <strong>Automatisation de migration de base de données :</strong> J'ai développé un script Bash pour automatiser entièrement la migration d'une base de données de <strong>MySQL 5 vers MySQL 8</strong>. Ce script intégrait des points de contrôle à chaque étape cruciale (sauvegarde, mise à jour, redémarrage du service) et générait un fichier de log détaillé pour assurer la traçabilité et la sécurité de l'opération.
        </li>
        <li>
            <strong>Gestion quotidienne de l'environnement :</strong> J'utilise couramment Bash pour naviguer dans le système de fichiers, gérer mes projets avec Git, manipuler des conteneurs Docker et lancer mes applications.
        </li>
    </ul>
</section>
<section class="comp-section">
    <h3>Projets Associés</h3>
    <div class="project-list">
        <ul>
            <li><a class="project-link" href="/projets/detail.html?id=1">Escarcelle</a></li>
        </ul>
    </div>
</section>
`
