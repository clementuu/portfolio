import { describe, it, expect, vi, afterEach } from 'vitest';
import { screen } from '@testing-library/dom';
import { render } from '@testing-library/svelte';

// Mock child components
vi.mock('../elements/OnScrollAppear.svelte', async () => {
    const mod = await import('../tests/mock/onScrollAppearMock.svelte');
    return { default: mod.default };
});
vi.mock('../cv/formationCard.svelte', async () => {
    const mod = await import('../tests/mock/formationCardMock.svelte');
    return { default: mod.default };
});
vi.mock('../cv/experienceCard.svelte', async () => {
    const mod = await import('../tests/mock/experienceCardMock.svelte');
    return { default: mod.default };
});

describe('CV component', () => {
    afterEach(() => {
        vi.restoreAllMocks();
        document.body.innerHTML = '';
    });

    it('renders "Chargement..." messages when no data is returned initially', async () => {
        // Stub fetch to return empty arrays immediately
        vi.stubGlobal('fetch', vi.fn((url) => {
            if (url.includes('/formations') || url.includes('/experiences')) {
                return Promise.resolve({
                    ok: true,
                    json: () => Promise.resolve([])
                });
            }
            return Promise.reject(new Error('unknown URL'));
        }));

        const CV = (await import('../cv/cv.svelte')).default;
        render(CV);

        // Expect to find the loading messages
        expect(await screen.findByText('Chargement des formations...')).toBeTruthy();
        expect(await screen.findByText('Chargement des expériences...')).toBeTruthy();

        // After data is loaded (even if empty), headings should appear
        expect(await screen.findByText('Formations')).toBeTruthy();
        expect(await screen.findByText('Expériences professionnelles')).toBeTruthy();
    });
});
