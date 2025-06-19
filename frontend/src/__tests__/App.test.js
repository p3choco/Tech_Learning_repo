import { render, screen, waitFor } from '@testing-library/react';
import axios from 'axios';
import App from '../App';

jest.mock('axios');

test('shows items list got from api', async () => {
  axios.get.mockResolvedValueOnce({
    data: [
      { id: 1, name: 'Item 1', description: 'Description 1' },
      { id: 2, name: 'Item 2', description: 'Description 2' },
    ],
  });

  render(<App />);

  await waitFor(() => expect(screen.queryByText(/Loading/)).toBeNull());

  const cards = screen.getAllByRole('heading', { level: 2 });
  expect(cards).toHaveLength(2);
  expect(cards[0]).toHaveTextContent('Item 1');
});
