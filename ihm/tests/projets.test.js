import { describe, it, expect, vi, afterEach } from 'vitest';
import { screen } from '@testing-library/dom';
import { render } from '@testing-library/svelte';

// Mock Card AVANT d'importer Projets
vi.mock('./mock/projetCardMock.svelte', async () => {
  const mod = await import('./mock/projetCardMock.svelte');
  return { default: mod.default };
});
// Redirige les imports de card.svelte vers notre mock
vi.mock('../projets/card.svelte', async () => {
  const mod = await import('./mock/projetCardMock.svelte');
  return { default: mod.default };
});

describe('Projets - simple tests (mock fetch before import)', () => {
  afterEach(() => {
    vi.restoreAllMocks();
    document.body.innerHTML = '';
  });

  it('renders projects when API returns data', async () => {
    const fakeProjets = [
      { id: 1, title: 'Alpha' },
      { id: 2, title: 'Beta' }
    ];

    // stub fetch AVANT d'importer le composant
    vi.stubGlobal('fetch', vi.fn(() =>
      Promise.resolve({
        ok: true,
        json: () => Promise.resolve(fakeProjets)
      })
    ));

    // importer après le stub (dynamic import)
    const Projets = (await import('../projets/projets.svelte')).default;

    render(Projets);

    // findByText attend l'apparition asynchrone
    expect(await screen.findByText('Mes Projets')).toBeTruthy();
    expect(await screen.findByText('Alpha')).toBeTruthy();
    expect(await screen.findByText('Beta')).toBeTruthy();
  });

  it('renders nothing when API returns empty array', async () => {
    vi.stubGlobal('fetch', vi.fn(() =>
      Promise.resolve({
        ok: true,
        json: () => Promise.resolve([])
      })
    ));

    const Projets = (await import('../projets/projets.svelte')).default;
    render(Projets);

    const maybeTitle = screen.queryByText('Mes Projets');
    expect(maybeTitle).toBeNull();
  });

  it('renders error message when fetch fails', async () => {
    vi.stubGlobal('fetch', vi.fn(() =>
      Promise.resolve({
        ok: false,
        status: 500,
        statusText: 'Server Error'
      })
    ));

    const Projets = (await import('../projets/projets.svelte')).default;
    render(Projets);

    expect(await screen.findByText('Erreur de la récupération des compétences.')).toBeTruthy();
  });
});
