import { render, screen } from '@testing-library/svelte';
import { describe, it, expect } from 'vitest';
import App from '../App.svelte';

describe('App', () => {
  it('should render the name "Clément Calia"', () => {
    render(App);
    expect(screen.getByText('Clément Calia')).toBeInTheDocument();
  });
});
