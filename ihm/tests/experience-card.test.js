import { describe, it, expect, afterEach } from 'vitest';
import { screen } from '@testing-library/dom';
import { render } from '@testing-library/svelte';
import ExperienceCard from '../cv/experienceCard.svelte';
import TestWrapper from './TestWrapper.svelte';

describe('ExperienceCard component', () => {
    afterEach(() => {
        document.body.innerHTML = '';
    });

    const sampleExperience = {
        Intitule: 'Développeur Full Stack',
        Type: 'CDI',
        Structure: 'Tech Solutions Inc.',
        Periode: 'Jan 2022 - Présent',
        Taches: [
            'Développement de nouvelles fonctionnalités',
            'Maintenance corrective et évolutive',
            'Participation aux revues de code'
        ]
    };

    it('renders experience details correctly', () => {
        render(TestWrapper, { props: { Component: ExperienceCard, props: { experience: sampleExperience } } });

        expect(screen.getByText('Développeur Full Stack [CDI]')).toBeInTheDocument();
        expect(screen.getByText('Tech Solutions Inc.')).toBeInTheDocument();
        expect(screen.getByText('| Jan 2022 - Présent')).toBeInTheDocument();
        expect(screen.getByText('Développement de nouvelles fonctionnalités')).toBeInTheDocument();
        expect(screen.getByText('Maintenance corrective et évolutive')).toBeInTheDocument();
        expect(screen.getByText('Participation aux revues de code')).toBeInTheDocument();
    });
});
