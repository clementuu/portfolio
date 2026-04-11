<svelte:options customElement="about-portfolio" />

<script>
  import Header from "./elements/header.svelte";

</script>

<Header/>

<div class="article">
    <section class="container about">
        <h1><b>À propos</b></h1>
        <article>
            <h4>Ce qui m'anime</h4>
            <p>
                Mon premier contact avec l'informatique remonte à l'enfance, j'était le petit-fils que les grands-parents appelaient
                à la rescousse quand l'ordinateur faisait des siennes, J'allais bidoullier un peu dans le panneau de navigation 
                et c'était très souvent c'était suffisant.
                <br>J'ai toujours été un vrai « touche à tout ». J'ai toujours eu besoin d'<b>explorer</b>, de <b>comprendre</b>, 
                de <b>tester</b>. Cette curiosité m'a suivi partout : dans mes lectures, mes projets, mes passions… et a forgé ma vision actuelle : 
                l'ingénierie ne se résume pas à la technique, elle prend tout son sens lorsqu'elle se met au service des personnes.
            </p>
            <p>
                C'est ce que j'aime appeller l'<b>innovation responsable</b> : créer des solutions durables, inclusives, qui s'adaptent aux réalités 
                du terrain pour servir l'intérêt général et réduire les fractures sociales.
            </p>

            <h4>Comment j'en suis arrivé là</h4>
            <p>
                Ma première ligne de code remonte à ma première année de licence. 
                Pour être honnête, ce n'était pas le coup de foudre immédiat. J'ai même eu quelques embrouilles avec 
                le compilateur Java, qui semblait prendre un malin plaisir à me renvoyer des erreurs incompréhensibles. 
                Mais j'ai rapidement découvert le plaisir du problème résolu, du « ça y est, ça marche ». 
                C'est là que l'informatique a commencé à prendre une vraie place.
            </p>
            <p>
                Après ma licence de mathématiques et informatique appliquées aux sciences humaines et sociales, mon service 
                civique m'a fait découvrir la gestion de projet et l'impact concret qu'un outil bien pensé peut avoir. 
                C'est ce qui m'a poussé à poursuivre en master d'expert en ingénierie logicielle, avec l'envie de créer 
                des applications <b>utiles</b>, <b>accessibles</b> et <b>robustes</b>.
            </p>
            <p>
                Depuis, j'ai travaillé sur des projets variés, me premettant d'acquérir une certaine <b>polyvalence</b>. Notamment la création d'un jeu éducatif, 
                l'<b>automatisation de tests End-2-End</b> d'un logiciel utilisant la reconnaissance faciale, la <b>refonte d'une application</b> de caisse, 
                la <b>modernisation d'un SI</b> pour des épiceries solidaires ou encore une <b>analyse statistique</b> du gaspillage alimentaire.
                <br>Ces expériences renforcent mon goût pour les solutions simples, maintenables et tournées vers l'humain.
            </p>

            <h4>Ce que j'apporte</h4>
            <p>
                La <b>curiosité</b>, la <b>rigueur</b> et le <b>sens du partage</b> sont au cœur de ma manière de travailler. J'aime comprendre 
                en profondeur, documenter ce que j'apprends, structurer ce que je fais et faciliter la vie de ceux qui 
                passeront derrière moi. 
                Douze années de sports collectifs - rugby puis handball - m'ont appris l'importance du <b>collectif</b>, de 
                <b>l'écoute</b> et de <b>l'humilité</b>. Ces réflexes se retrouvent aujourd'hui dans ma façon de collaborer, d'animer 
                des réunions ou de préparer des présentations professionnelles.
            </p>
            <p>
                J'apporte aujourd'hui une <b>approche globale</b>, faisant le lien entre les enjeux métiers et les contraintes techniques. 
                Grâce à mon <b>esprit critique</b>, je cherche à anticiper les risques pour proposer des outils véritablement utiles, 
                conçus pour durer et faciliter la vie de tous les utilisateurs.
            </p>

            <h4>Ce qui me nourrit</h4>
            <p>
                En dehors du code, je me ressource dans la <b>musique</b> : je joue de la guitare depuis le collège et je produis 
                un peu quand j'ai le temps. Les <b>sciences</b> et la <b>culture</b> alimentent ma curiosité, et le <b>sport</b> m'aide à garder 
                un bon équilibre. 
                Ces centres d'intérêt ne sont pas des activités à côté : ils influencent ma manière de penser, de créer, 
                d'aborder les problèmes. Ils m'aident à garder les idées fraîches et à relier des univers qui, au premier 
                abord, n'ont rien à voir.
                Ayant grandi à la campagne, la <b>nature</b> occupe une place essentielle et nourrit ma conscience écologique, 
                m'orientant naturellement vers l'<b>écoconception</b> numérique.
            </p>
            <p>
                Ma curiosité me pousse aussi vers des projets plus personnels, comme le développement d'un VST pour explorer la synthèse 
                de son ou la contribution à l'<b>open-source</b>. Pour moi, l'innovation doit rester profondément humaine, créative et responsable.
            </p>
        </article>
    </section>
</div>

<style>
    .article {
        position: relative;
        top: 0;
        margin: auto;
    }

    .about {  
        max-width: var(--mobile-width);
        margin: 0 auto;
        padding: 2em;
        line-height: 1.6;
    }

    h1 {
        padding: 10px 0;
        margin-bottom:1em;
    }

    h4 {
        color: var(--secondary-color);
        margin: 1em 0;
        padding-left: 0;
    }

    @media (min-width: 800px) {
		.article {
			top: var(--header-height);
            padding: 1em;
            margin: auto auto 5em auto;
		}
	}

</style>
