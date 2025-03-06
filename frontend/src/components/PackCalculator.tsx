import React, { useState } from 'react';

const PacksCalculator: React.FC = () => {
    const apiUrl = import.meta.env.VITE_API_URL;
    const [order, setOrder] = useState<number>(0);
    const [packs, setPacks] = useState<{ size: number, count: number }[]>([]);
    const [error, setError] = useState<string>('');

    const handleSubmit = async (e: React.FormEvent) => {
        e.preventDefault();
        if (order <= 0) {
            setError('Order quantity must be a positive integer.');
            return;
        }

        try {
            const response = await fetch(`${apiUrl}/calculate?order=${order}`);
            if (!response.ok) {
                throw new Error('Failed to fetch pack calculation.');
            }
            const data = await response.json();
            setPacks(data);
            setError('');
        } catch (err) {
            if (err instanceof Error) {
                setError(err.message || 'An error occurred while calculating packs.');
            } else {
                setError('An error occurred while calculating packs.');
            }
        }
    };

    return (
        <div>
            <h1>Packs Calculator</h1>
            <form onSubmit={handleSubmit}>
                <div>
                    <label htmlFor="order">Enter Order Quantity:</label>
                    <input
                        type="number"
                        id="order"
                        value={order}
                        onChange={(e) => setOrder(Number(e.target.value))}
                        min={1}
                        required
                    />
                </div>
                <button type="submit">Calculate Packs</button>
            </form>

            {error && <p style={{ color: 'red' }}>{error}</p>}

            {packs.length > 0 && (
                <div>
                    <h2>Optimal Packs</h2>
                    <ul>
                        {packs.map((pack, index) => (
                            <li key={index}>
                                {pack.count} x {pack.size} items
                            </li>
                        ))}
                    </ul>
                </div>
            )}
        </div>
    );
};

export default PacksCalculator;
