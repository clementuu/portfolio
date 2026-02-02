import { describe, it, expect, afterEach } from 'vitest';
import { screen } from '@testing-library/dom';
import { render } from '@testing-library/svelte';
import FormationCard from '../cv/formationCard.svelte';

describe('FormationCard component', () => {
    afterEach(() => {
        document.body.innerHTML = '';
    });

    const sampleFormation = {
        Intitule: 'Master Ingénierie Logicielle',
        Etablissement: 'Université de Montpellier',
        Periode: '2020 - 2022'
    };

    it('renders formation details correctly', () => {
        render(FormationCard, { props: { formation: sampleFormation } });

        expect(screen.getByText('Master Ingénierie Logicielle')).toBeInTheDocument();
        expect(screen.getByText('Université de Montpellier')).toBeInTheDocument();
        expect(screen.getByText('| 2020 - 2022')).toBeInTheDocument();
    });
});
