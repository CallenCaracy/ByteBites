import React, { useState } from 'react';
import { useMutation } from '@apollo/client';
import { CREATE_INVENTORY } from '../graphql/Kitchenqueries';
import Navbar from '../components/NavBar';
import { useNavigate, useParams } from 'react-router-dom';

const MakeInventory = () => {
  const navigate = useNavigate();
  const { menuId } = useParams();

  const [formData, setFormData] = useState({
    menuID: menuId || '',
    availableServings: '',
    lowStockThreshold: ''
  });

  const [createInventory, { data, loading, error }] = useMutation(CREATE_INVENTORY);

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const { name, value } = e.target;
    setFormData(prev => ({
      ...prev,
      [name]: value,
    }));
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();

    try {
      await createInventory({
        variables: {
          menuId: formData.menuID,
          availableServings: parseInt(formData.availableServings),
          lowStockThreshold: formData.lowStockThreshold ? parseInt(formData.lowStockThreshold) : undefined
        }
      });

      navigate(`/dashboard`);
    } catch (err) {
      console.error("Error creating inventory:", err);
    }

    setFormData({
      menuID: menuId || '',
      availableServings: '',
      lowStockThreshold: ''
    });
  };

  return (
    <>
      <Navbar />
      <form
        onSubmit={handleSubmit}
        className="max-w-xl mx-auto p-4 shadow-md rounded-xl bg-white space-y-4 text-black mt-4"
      >
        <h2 className="text-xl font-bold">Add Inventory</h2>

        <div className="flex flex-col space-y-1">
          <label htmlFor="availableServings" className="font-medium">Available Servings</label>
          <input
            id="availableServings"
            name="availableServings"
            type="number"
            value={formData.availableServings}
            onChange={handleChange}
            className="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md"
            placeholder="e.g. 40"
            required
          />
        </div>

        <div className="flex flex-col space-y-1">
          <label htmlFor="lowStockThreshold" className="font-medium">Low Stock Threshold (Optional)</label>
          <input
            id="lowStockThreshold"
            name="lowStockThreshold"
            type="number"
            value={formData.lowStockThreshold}
            onChange={handleChange}
            className="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md"
            placeholder="e.g. 10"
            required
          />
        </div>

        <div className="flex space-x-2">
          <button
            type="submit"
            disabled={loading}
            className="btn btn-primary w-1/2 bg-blue-600 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded"
          >
            {loading ? "Adding..." : "Add Inventory"}
          </button>
        </div>

        {error && <p className="text-red-500 text-sm">Error: {error.message}</p>}
        {data && (
          <p className="text-green-600 text-sm">
            Inventory created with ID: {data.createInventory.id}
          </p>
        )}
      </form>
    </>
  );
};

export default MakeInventory;
