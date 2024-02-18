import { useEffect, useState } from "react";

const Navbar = () => {
  const [isScrolled, setIsScrolled] = useState(false);
  const [isMenuOpen, setIsMenuOpen] = useState(false);

  useEffect(() => {
    const handleScroll = () => {
      const scrollTop = window.scrollY;
      setIsScrolled(scrollTop > 0);
    };

    window.addEventListener("scroll", handleScroll);

    return () => {
      window.removeEventListener("scroll", handleScroll);
    };
  }, []);

  const handleMenu = () => {
    setIsMenuOpen(!isMenuOpen);
    console.log("diklik");
  };

  return (
    <div className={`flex items-center justify-center w-full ${isScrolled ? 'bg-white bg-opacity-80 fixed top-0 left-0 z-50 navbar-fixed' : 'bg-transparent absolute top-0 left-0 z-10'}`}>
      <div className="container">
        <div className="flex items-center py-4 justify-between relative">
          <div className="flex items-center px-4">
            <button
              id="hamburger"
              name="hamburger"
              className="block absolute left-4"
              onClick={handleMenu}
            >
              <span
                className={
                  isMenuOpen
                    ? "hamburger-active"
                    : "hamburger-line origin-top-left"
                }
              ></span>
              <span
                className={isMenuOpen ? "hamburger-active" : "hamburger-line"}
              ></span>
              <span
                className={
                  isMenuOpen
                    ? "hamburger-active "
                    : "hamburger-line origin-bottom-left"
                }
              ></span>
            </button>
          </div>
          <div className="px-4">
            <a href="#home" className="font-bold text-lg text-black">
              Jobseekers
            </a>
          </div>
        </div>
      </div>
    </div>
  );
};

export default Navbar;
