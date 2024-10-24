import React, { useState } from 'react';
import Navbar from './components/Navbar';

interface Burger {
  id: number;
  name: string;
  description: string;
  price: number;
  image: string;
}

const burgers: Burger[] = [
  {
    id: 1,
    name: 'Classic Burger',
    description: 'A juicy beef patty with lettuce, tomato, and cheese.',
    price: 8.99,
    image: 'classic-burger.jpg',
  },
  {
    id: 2,
    name: 'Cheese Lover’s Burger',
    description: 'A double cheeseburger for cheese enthusiasts.',
    price: 10.99,
    image: 'cheese-burger.jpg',
  },
  {
    id: 3,
    name: 'Veggie Burger',
    description: 'A vegetarian delight with fresh veggies and avocado.',
    price: 7.99,
    image: 'veggie-burger.jpg',
  },
];

const OrderingPage: React.FC = () => {
  const [cart, setCart] = useState<Burger[]>([]);

  const addToCart = (burger: Burger) => {
    setCart([...cart, burger]);
  };

  return (
    <div className="min-h-screen bg-gray-100">
      {/* Navbar */}
      <Navbar />

      {/* Ordering Page Content */}
      <div className="container mx-auto px-4 py-8">
        <h1 className="text-4xl font-bold mb-8 text-center">Order Your Favorite Burger</h1>
        
        {/* Burger List */}
        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-8">
          {burgers.map((burger) => (
            <div key={burger.id} className="bg-white shadow-md rounded-lg p-4">
              <img src={burger.image} alt={burger.name} className="h-48 w-full object-cover rounded-lg mb-4" />
              <h2 className="text-xl font-bold mb-2">{burger.name}</h2>
              <p className="text-gray-600 mb-4">{burger.description}</p>
              <p className="text-lg font-bold text-red-600 mb-4">${burger.price.toFixed(2)}</p>
              <button
                className="bg-red-600 text-white px-4 py-2 rounded-lg hover:bg-red-700 transition"
                onClick={() => addToCart(burger)}
              >
                Add to Cart
              </button>
            </div>
          ))}
        </div>
        
        {/* Cart Summary */}
        <div className="fixed bottom-4 right-4 bg-white p-6 rounded-lg shadow-lg">
          <h3 className="text-2xl font-bold mb-4">Cart Summary</h3>
          {cart.length > 0 ? (
            <div>
              {cart.map((item, index) => (
                <div key={index} className="mb-2">
                  <p>
                    {item.name} - ${item.price.toFixed(2)}
                  </p>
                </div>
              ))}
              <div className="font-bold text-lg">
                Total: ${cart.reduce((total, item) => total + item.price, 0).toFixed(2)}
              </div>
              <button
                className="bg-green-600 text-white mt-4 px-4 py-2 rounded-lg hover:bg-green-700 transition"
                onClick={() => alert('Proceed to checkout!')}
              >
                Checkout
              </button>
            </div>
          ) : (
            <p>Your cart is empty.</p>
          )}
        </div>
      </div>
    </div>
  );
};

export default OrderingPage;
