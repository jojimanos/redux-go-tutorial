import React from 'react';
import { NavLink, useNavigate, useOutletContext } from 'react-router-dom';
import { useEffect } from 'react';

const LandingPage: React.FC = () => {

    const {token} = useOutletContext<{token: string}>();
    const navigate = useNavigate();

    useEffect(() => {
    if (token)
        navigate("/profile");
    }, [token, navigate])

  return (
    <div className="bg-gray-100 min-h-screen flex flex-col justify-between">
      {/* Hero Section */}
      <div className="flex-grow">
        <header className="bg-white shadow">
          <nav className="container mx-auto flex justify-between items-center p-4">
            <div className="text-2xl font-bold text-red-600">BurgerMania</div>
            <ul className="flex space-x-4">
              <li>
                <NavLink to="/login" className="text-gray-600 hover:text-red-600">
                  Login
                </NavLink>
              </li>
            </ul>
          </nav>
        </header>

        <main className="container mx-auto flex flex-col items-center justify-center flex-grow px-4 py-16 text-center">
          <h1 className="text-5xl font-bold text-gray-800 mb-4">
            Welcome to BurgerMania
          </h1>
          <p className="text-lg text-gray-600 mb-8">
            Craving a delicious burger? Order now and enjoy the best burger in town!
          </p>
          <NavLink to="/login">
            <button className="bg-red-600 text-white px-6 py-3 rounded-full text-lg font-semibold hover:bg-red-700 transition duration-300">
              Get Started
            </button>
          </NavLink>
        </main>
      </div>

      {/* Footer */}
      <footer className="bg-gray-800 text-white py-4">
        <div className="container mx-auto text-center">
          <p>© 2024 BurgerMania. All Rights Reserved.</p>
        </div>
      </footer>
    </div>
  );
};

export default LandingPage;
