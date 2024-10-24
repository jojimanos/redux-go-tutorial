import React from 'react';
import { Link, useNavigate } from 'react-router-dom';

const Navbar: React.FC = () => {
  const navigate = useNavigate();

  const handleLogout = () => {
    // Clear token from localStorage and navigate to login page
    localStorage.removeItem('token');
    navigate('/login');
  };

  return (
    <nav className="bg-gray-800 text-white p-4">
      <div className="container mx-auto flex justify-between items-center">
        <Link to="/" className="text-2xl font-bold">
          BurgerMania
        </Link>
        <ul className="flex space-x-4">
          <li>
            <Link to="/" className="hover:text-red-500">
              Main Page
            </Link>
          </li>
          <li>
            <Link to="/profile" className="hover:text-red-500">
              Profile
            </Link>
          </li>
          <li>
            <button onClick={handleLogout} className="hover:text-red-500">
              Logout
            </button>
          </li>
        </ul>
        <Link to="/cart" className="flex items-center">
          <span className="material-icons">shopping_cart</span>
          <span className="ml-1">(3)</span> {/* This is the cart item count */}
        </Link>
      </div>
    </nav>
  );
};

export default Navbar;
