import { describe, it, expect, vi, afterEach } from 'vitest';
import { screen } from '@testing-library/dom';
import { render } from '@testing-library/svelte';

// Mock Card AVANT d'importer Competences
vi.mock('./mock/competenceCardMock.svelte', async () => {
  const mod = await import('./mock/competenceCardMock.svelte');
  return { default: mod.default };
});
// Redirige les imports de card.svelte vers notre mock
vi.mock('../competences/card.svelte', async () => {
  const mod = await import('./mock/competenceCardMock.svelte');
  return { default: mod.default };
});

describe('Competences - minimal tests', () => {
  afterEach(() => {
    vi.restoreAllMocks();
    document.body.innerHTML = '';
  });

  it('renders competences when API returns data', async () => {
    const fakeCompetences = [
      { id: 1, title: 'Go', type: 'backend' },
      { id: 2, title: 'Svelte', type: 'frontend' }
    ];

    vi.stubGlobal('fetch', vi.fn(() =>
      Promise.resolve({
        ok: true,
        json: () => Promise.resolve(fakeCompetences)
      })
    ));

    const Competences = (await import('../competences/competences.svelte')).default;
    render(Competences);

    expect(await screen.findByText('Mes Compétences')).toBeTruthy();
    expect(await screen.findByText('Mock Card: Go')).toBeTruthy();
    expect(await screen.findByText('Mock Card: Svelte')).toBeTruthy();
  });

  it('renders nothing when API returns empty array', async () => {
    vi.stubGlobal('fetch', vi.fn(() =>
      Promise.resolve({
        ok: true,
        json: () => Promise.resolve([])
      })
    ));

    const Competences = (await import('../competences/competences.svelte')).default;
    render(Competences);

    const maybeTitle = screen.queryByText('Mes Compétences');
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

    const Competences = (await import('../competences/competences.svelte')).default;
    render(Competences);

    expect(await screen.findByText('Erreur lors du chargement des compétences.')).toBeTruthy();
  });

  it('filters competences by type from URL query parameter', async () => {
    const fakeCompetences = [
      { id: 1, title: 'Go', type: 'backend' },
      { id: 2, title: 'Svelte', type: 'frontend' },
      { id: 3, title: 'Docker', type: 'devops' }
    ];

    vi.stubGlobal('fetch', vi.fn(() =>
      Promise.resolve({
        ok: true,
        json: () => Promise.resolve(fakeCompetences)
      })
    ));

    // Mock window.location.search for URL parameter testing
    Object.defineProperty(window, 'location', {
        value: {
            search: '?type=backend',
        },
        writable: true,
    });

    const Competences = (await import('../competences/competences.svelte')).default;
    render(Competences);

    expect(await screen.findByText('Mes Compétences')).toBeTruthy();
    expect(await screen.findByText('Mock Card: Go')).toBeTruthy();
    expect(screen.queryByText('Mock Card: Svelte')).toBeNull();
    expect(screen.queryByText('Mock Card: Docker')).toBeNull();
  });
});
