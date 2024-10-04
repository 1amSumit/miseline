import { useState } from "react";
import { FaEye, FaEyeSlash, FaRegEye } from "react-icons/fa";

export default function Login() {
  const [showPassword, setShowPassword] = useState(false);
  return (
    <div className=" h-screen w-screen flex items-center  justify-center">
      <div className="grid grid-cols-2 gap-4 w-[70%] h-[70%] border-[1px] border-gray-200 rounded-lg shadow-lg">
        <figure className="">
          <img
            className="w-[100%] h-[100%] rounded-tl-lg rounded-bl-lg"
            src="/shop-2.jpg"
            loading="lazy"
          />
        </figure>
        <div className="flex flex-col items-center px-2 py-3">
          <figure className="w-[20%] h-[20%] flex flex-col items-center">
            <img src="/logo.png" className="w-full h-full" />
            <div>
              <p className="font-medium font-roboto text-gray-700">MiseLink</p>
            </div>
          </figure>
          <div className="mt-[2rem] flex flex-col text-gray-800 items-center">
            <p className="font-roboto text-lg font-bold">
              Sign in to your Account
            </p>
            <p className="text-sm text-gray-500">
              Enter you email and password
            </p>
          </div>
          <div className="mt-[2rem]  flex flex-col gap-2">
            <div className="flex flex-row gap-4 border-[1px] py-2 rounded-lg border-gray-300 ">
              <input
                type="email"
                className="placeholder:text-sm px-3 shadow-sm outline-none"
                placeholder="Enter your email"
              />
            </div>
            <div className="flex flex-row gap-4 border-[1px] py-2 rounded-lg border-gray-300 items-center justify-center px-2">
              <input
                type={showPassword ? "text" : "password"}
                className=" placeholder:text-sm px-3 shadow-sm outline-none"
                placeholder="Enter your password"
              />
              {showPassword ? (
                <FaEye
                  className="cursor-pointer"
                  onClick={() => {
                    setShowPassword(false);
                  }}
                />
              ) : (
                <FaEyeSlash
                  className="cursor-pointer"
                  onClick={() => {
                    setShowPassword(true);
                  }}
                />
              )}
            </div>

            <div>
              <button className="bg-[#007CC1] text-white w-full py-1 rounded-lg hover:bg-[#35a6e3] transition-all duration-300">
                Login
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}
