import React, { useState, useEffect } from 'react';
import Navbar from './components/Navbar';

interface Order {
  id: number;
  burger: string;
  date: string;
  total: number;
}

const ProfilePage: React.FC = () => {
  const [user, setUser] = useState({
    name: 'John Doe',
    email: 'johndoe@example.com',
  });

  const [orders, setOrders] = useState<Order[]>([
    {
      id: 1,
      burger: 'Classic Burger',
      date: '2024-10-20',
      total: 8.99,
    },
    {
      id: 2,
      burger: 'Veggie Burger',
      date: '2024-10-21',
      total: 7.99,
    },
  ]);

  const [editMode, setEditMode] = useState(false);
  const [name, setName] = useState(user.name);
  const [email, setEmail] = useState(user.email);

  const handleEdit = () => {
    setEditMode(!editMode);
  };

  const handleSave = () => {
    setUser({ name, email });
    setEditMode(false);
  };

  return (
    <div className="min-h-screen bg-gray-100">
      {/* Navbar */}
      <Navbar />

      <div className="container mx-auto p-6">
        <h1 className="text-4xl font-bold mb-8 text-center">Profile Page</h1>
        
        {/* User Info Section */}
        <div className="bg-white shadow-md p-6 rounded-lg mb-8">
          <h2 className="text-2xl font-bold mb-4">User Information</h2>
          <div className="mb-4">
            <label className="block text-gray-600">Name:</label>
            {editMode ? (
              <input
                type="text"
                value={name}
                onChange={(e) => setName(e.target.value)}
                className="border p-2 rounded-lg w-full"
              />
            ) : (
              <p>{user.name}</p>
            )}
          </div>
          <div className="mb-4">
            <label className="block text-gray-600">Email:</label>
            {editMode ? (
              <input
                type="email"
                value={email}
                onChange={(e) => setEmail(e.target.value)}
                className="border p-2 rounded-lg w-full"
              />
            ) : (
              <p>{user.email}</p>
            )}
          </div>

          {editMode ? (
            <div>
              <button
                className="bg-green-600 text-white px-4 py-2 rounded-lg mr-2"
                onClick={handleSave}
              >
                Save
              </button>
              <button
                className="bg-red-600 text-white px-4 py-2 rounded-lg"
                onClick={handleEdit}
              >
                Cancel
              </button>
            </div>
          ) : (
            <button
              className="bg-blue-600 text-white px-4 py-2 rounded-lg"
              onClick={handleEdit}
            >
              Edit Profile
            </button>
          )}
        </div>

        {/* Order History Section */}
        <div className="bg-white shadow-md p-6 rounded-lg">
          <h2 className="text-2xl font-bold mb-4">Order History</h2>
          {orders.length > 0 ? (
            <table className="table-auto w-full text-left">
              <thead>
                <tr>
                  <th className="px-4 py-2">Order ID</th>
                  <th className="px-4 py-2">Burger</th>
                  <th className="px-4 py-2">Date</th>
                  <th className="px-4 py-2">Total</th>
                </tr>
              </thead>
              <tbody>
                {orders.map((order) => (
                  <tr key={order.id} className="border-t">
                    <td className="px-4 py-2">{order.id}</td>
                    <td className="px-4 py-2">{order.burger}</td>
                    <td className="px-4 py-2">{order.date}</td>
                    <td className="px-4 py-2">${order.total.toFixed(2)}</td>
                  </tr>
                ))}
              </tbody>
            </table>
          ) : (
            <p>No orders yet.</p>
          )}
        </div>
      </div>
    </div>
  );
};

export default ProfilePage;
