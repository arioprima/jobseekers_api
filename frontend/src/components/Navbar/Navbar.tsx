import { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";

interface Link {
  name: string;
  url: string;
}

const Navbar = () => {
  const [isScrolled, setIsScrolled] = useState(false);
  const [isMenuOpen, setIsMenuOpen] = useState(false);
  const navigate = useNavigate();

  const links: Link[] = [
    { name: "Tentang", url: "/tentang" },
    { name: "Mitra", url: "/mitra" },
    { name: "Hubungi", url: "/hubungi" },
  ];

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
    <div
      className={`flex items-center justify-center w-full ${
        isScrolled
          ? "bg-white bg-opacity-80 fixed top-0 left-0 z-50 navbar-fixed"
          : "bg-transparent absolute top-0 left-0 z-10"
      }`}
    >
      <div className="container">
        <div className="flex items-center py-1 justify-between relative">
          <div className="flex flex-row py-4 lg:flex-row-reverse lg:justify-between lg:py-2 w-full justify-center">
            <div className="flex items-center px-4">
              <button
                id="hamburger"
                name="hamburger"
                className="block absolute left-4 lg:hidden"
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
              <nav
                id="nav-menu"
                className={
                  isMenuOpen
                    ? "bg-white absolute py-5 shadow-lg rounded-lg max-w-[150px] w-full left-4 top-full"
                    : "hidden lg:block lg:static bg-transparent lg:max-w-full lg:shadow-none lg:rounded-none"
                }
              >
                <ul className="block lg:flex lg:gap-5">
                  {links.map((link, index) => (
                    <li key={index} className="group">
                      <a
                       onClick={() => navigate(link.url)}
                        className="text-base text-black py-2 flex group-hover:text-sky-500 hover:cursor-pointer"
                      >
                        {link.name}
                      </a>
                    </li>
                  ))}
                </ul>
              </nav>
            </div>
            <div className="pl-8 lg:pt-1.5">
              <a onClick={() => navigate("/")} className="font-bold text-2xl text-black hover:cursor-pointer">
                Jobseekers
              </a>
            </div>
          </div>
          <div className="pl-4 pt-1 lg:pt-0 hover:cursor-pointer">
            <a
                onClick={() => navigate("/login")}
              className="text-base font-semibold text-md text-white bg-green-500 py-3 px-3 rounded-xl
                hover:shadow-lg hover:opacity-85 transition duration-300 ease-in-out
                "
            >
              Masuk
            </a>
          </div>
        </div>
      </div>
    </div>
  );
};

export default Navbar;
