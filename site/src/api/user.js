import $axios from "./index";

//用户相关API

function login(data) {
  const url = "/base/login";
  return $axios.post(url, data);
}

function register(data) {
  const url = "/base/register";
  return $axios.post(url, data);
}

export default {
  login,
  register,
};
