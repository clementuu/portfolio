import { describe, it, expect, vi, afterEach } from 'vitest';
import { render, screen } from '@testing-library/svelte';

// IMPORTANT : on doit stubber fetch et préparer window.location.search
// AVANT d'importer le composant si le fetch est exécuté au niveau du module.

describe('ProjetDetail component (simple, reliable tests)', () => {
  afterEach(() => {
    vi.restoreAllMocks();
    // reset location to root between tests
    history.pushState({}, '', '/');
    document.body.innerHTML = '';
  });

  it("shows an error when no id is provided in the URL", async () => {
    // ensure URL has no id
    history.pushState({}, '', '/');

    // import AFTER having set the URL (no fetch stub needed)
    const ProjetDetail = (await import('../projets/detail.svelte')).default;

    render(ProjetDetail);

    // le composant définit error = "Aucun ID de projet n'a été fourni."
    expect(await screen.findByText("Aucun ID de projet n'a été fourni.")).toBeTruthy();
  });

  it('renders project details when API returns a projet', async () => {
    // set URL with id query param
    history.pushState({}, '', '/?id=42');

    // fake projet data with tmpl containing HTML we can assert
    const fakeProjet = {
      id: 42,
      name: 'Projet Zeta',
      sujet: 'Sujet important',
      tmpl: '<div class="project-section"><h2>Présentation</h2><p>Contenu</p></div>'
    };

    // stub fetch BEFORE importing the component
    vi.stubGlobal('fetch', vi.fn(() =>
      Promise.resolve({
        ok: true,
        json: () => Promise.resolve(fakeProjet)
      })
    ));

    const ProjetDetail = (await import('../projets/detail.svelte')).default;

    render(ProjetDetail);

    // on attend que le titre soit rendu
    expect(await screen.findByText('Projet Zeta')).toBeTruthy();
    expect(await screen.findByText('Sujet important')).toBeTruthy();

    // et que le HTML injecté via {@html} soit présent (on peut chercher le texte ou la balise)
    expect(await screen.findByText('Présentation')).toBeTruthy();
    expect(document.querySelector('.project-section')).toBeTruthy();
  });

  it('shows an error message when the fetch fails', async () => {
    // set URL with id
    history.pushState({}, '', '/?id=99');

    // stub fetch to return non-ok
    vi.stubGlobal('fetch', vi.fn(() =>
      Promise.resolve({
        ok: false,
        status: 500,
        statusText: 'Server Error'
      })
    ));

    const ProjetDetail = (await import('../projets/detail.svelte')).default;

    render(ProjetDetail);

    // le composant définit error = "Erreur lors du chargement des détails du projet."
    expect(await screen.findByText('Erreur lors du chargement des détails du projet.')).toBeTruthy();
  });
});
