import { describe, it, expect, vi, afterEach } from 'vitest';
import { render, screen } from '@testing-library/svelte';

describe('CompetenceDetail component', () => {
  afterEach(() => {
    vi.restoreAllMocks();
    history.pushState({}, '', '/'); // Reset URL
    document.body.innerHTML = '';
  });

  it("shows an error when no id is provided in the URL", async () => {
    history.pushState({}, '', '/'); // Ensure no id

    const CompetenceDetail = (await import('../competences/detail.svelte')).default;
    render(CompetenceDetail);

    expect(await screen.findByText("Aucun ID de compétence n'a été fourni.")).toBeTruthy();
  });

  it('renders competence details when API returns a competence', async () => {
    history.pushState({}, '', '/?id=1');

    const fakeCompetence = {
      id: 1,
      name: 'GoLang',
      image: 'path/to/go.png',
      rating: 4,
      template: '<div class="comp-section"><h2>Description</h2><p>Contenu GoLang</p></div>'
    };

    vi.stubGlobal('fetch', vi.fn(() =>
      Promise.resolve({
        ok: true,
        json: () => Promise.resolve(fakeCompetence)
      })
    ));

    const CompetenceDetail = (await import('../competences/detail.svelte')).default;
    render(CompetenceDetail);

    expect(await screen.findByText('GoLang')).toBeTruthy();
    expect(await screen.findByText('⭐'.repeat(4) + '☆'.repeat(1))).toBeTruthy(); // Check stars
    expect(await screen.findByText('Description')).toBeTruthy();
    expect(await screen.findByText('Contenu GoLang')).toBeTruthy();
    expect(document.querySelector('img[alt="Logo de GoLang"]')).toBeTruthy();
  });

  it('shows an error message when the fetch fails', async () => {
    history.pushState({}, '', '/?id=99');

    vi.stubGlobal('fetch', vi.fn(() =>
      Promise.resolve({
        ok: false,
        status: 500,
        statusText: 'Server Error'
      })
    ));

    const CompetenceDetail = (await import('../competences/detail.svelte')).default;
    render(CompetenceDetail);

    expect(await screen.findByText('Erreur lors du chargement des détails de la compétence.')).toBeTruthy();
  });
});
