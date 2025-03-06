import React, { useEffect, useState } from 'react';
import axios from 'axios';

interface Pack {
    id: number;
    size: number;
}

const PacksCrud: React.FC = () => {
    const apiUrl = import.meta.env.VITE_API_URL;
    const [packs, setPacks] = useState<Pack[]>([]);
    const [newPackSize, setNewPackSize] = useState<number>(0);
    const [updatedPackSize, setUpdatedPackSize] = useState<number>(0);
    const [updatePackId, setUpdatePackId] = useState<number | null>(null);
    const [error, setError] = useState<string>('');
    const [loading, setLoading] = useState<boolean>(false);

    // Fetch all packs
    const fetchPacks = async () => {
        setLoading(true);
        try {
            const response = await axios.get(`${apiUrl}/packs`);
            if (Array.isArray(response.data)) {
                setPacks(response.data);
            } else {
                setError('Failed to load packs: Unexpected data format');
            }
            setError('');
        } catch (err) {
            setError('Failed to fetch packs');
        } finally {
            setLoading(false);
        }
    };

    const handleAddPack = async () => {
        if (newPackSize <= 0) {
            setError('Pack size must be a positive number');
            return;
        }
        setError('');
        try {
            await axios.post(`${apiUrl}/packs`, { size: newPackSize });
            setNewPackSize(0);
            fetchPacks();
        } catch (err) {
            setError('Failed to add pack size');
        }
    };

    const handleUpdatePack = async () => {
        if (updatedPackSize <= 0 || updatePackId === null) {
            setError('Invalid input for update');
            return;
        }
        setError('');
        try {
            await axios.put(`${apiUrl}/packs/${updatePackId}`, { size: updatedPackSize });
            setUpdatedPackSize(0);
            setUpdatePackId(null);
            fetchPacks();
        } catch (err) {
            setError('Failed to update pack size');
        }
    };

    const handleDeletePack = async (id: number) => {
        setError('');
        try {
            await axios.delete(`${apiUrl}/packs/${id}`);
            fetchPacks();
        } catch (err) {
            setError('Failed to delete pack');
        }
    };

    useEffect(() => {
        fetchPacks();
    }, []);

    return (
        <div>
            <h2>Manage Pack Sizes</h2>

            {/* Add Pack Size */}
            <div>
                <h3>Add a Pack Size</h3>
                <input
                    type="number"
                    value={newPackSize}
                    onChange={(e) => setNewPackSize(Number(e.target.value))}
                    placeholder="Enter pack size"
                />
                <button onClick={handleAddPack} disabled={loading}>
                    {loading ? 'Adding...' : 'Add Pack'}
                </button>
            </div>

            {/* Update Pack Size */}
            <div>
                <h3>Update a Pack Size</h3>
                <select
                    value={updatePackId ?? ''}
                    onChange={(e) => setUpdatePackId(Number(e.target.value))}
                >
                    <option value="">Select Pack to Update</option>
                    {packs.length > 0 ? (
                        packs.map((pack) => (
                            <option key={pack.id} value={pack.id}>
                                {pack.size}
                            </option>
                        ))
                    ) : (
                        <option value="">No packs available</option>
                    )}
                </select>
                <input
                    type="number"
                    value={updatedPackSize}
                    onChange={(e) => setUpdatedPackSize(Number(e.target.value))}
                    placeholder="Enter new size"
                />
                <button onClick={handleUpdatePack} disabled={loading}>
                    {loading ? 'Updating...' : 'Update Pack'}
                </button>
            </div>

            {/* Display Packs */}
            <h3>Existing Pack Sizes</h3>
            {error && <p style={{ color: 'red' }}>{error}</p>}

            {loading ? (
                <p>Loading...</p>
            ) : packs.length === 0 ? (
                <p>No packs available.</p>
            ) : (
                <ul>
                    {packs.map((pack) => (
                        <li key={pack.id}>
                            {pack.size}{' '}
                            <button onClick={() => handleDeletePack(pack.id)} disabled={loading}>
                                {loading ? 'Deleting...' : 'Delete'}
                            </button>
                        </li>
                    ))}
                </ul>
            )}
        </div>
    );
};

export default PacksCrud;
