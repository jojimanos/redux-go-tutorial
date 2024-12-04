import React from 'react';
import { Link } from 'react-router-dom';
import { FaShoppingCart } from "react-icons/fa";
import { useDispatch } from 'react-redux';
import { userActions } from '../../stateSlices/userSlice';

const Navbar: React.FC = () => {
  const dispatch = useDispatch()

  const handleLogout = () => {
    // Clear token from localStorage and navigate to login page
    localStorage.removeItem('token');
    dispatch(userActions.logout())
  };

  return (
    <nav className="bg-gray-800 text-white p-4">
      <div className="container mx-auto flex justify-between items-center">
        <Link to="/" className="text-2xl font-bold">
          BurgerMania
        </Link>
        
        
        <ul className="flex space-x-4">
          <li>
            <Link to="/profile" className="hover:text-red-500">
              Profile
            </Link>
          </li>
          <li>
<Link to="/cart" className="flex items-center">
          <span className="material-icons">
<FaShoppingCart />
          </span>
          <span className="ml-1">(3)</span> {/* This is the cart item count */}
</Link>
          </li>
          <li>
            <button onClick={handleLogout} className="hover:text-red-500">
              Logout
            </button>
          </li>
        </ul>
      </div>
    </nav>
  );
};

export default Navbar;
