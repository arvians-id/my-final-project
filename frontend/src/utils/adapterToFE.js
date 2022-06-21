/**
 * Collection of function to adapting data from Backend to Frontend
 * useful when we want to GET some data from BE
 */

export const adapterUserToFE = (user) => {
  return {
    id: user.id,
    name: user.name,
    username: user.username,
    email: user.email,
    role: user.role === 1 ? 'Guru' : user.role === 2 ? 'Siswa' : 'Admin',
    gender: user.gender === 1 ? "Pria" : user.gender === 2 ? "Wanita" : "Lainnya",
    type_of_disability: user.type_of_disability === 1 ? "Tunanetra" : user.type_of_disability === 2 ? "Tunarungu" : "-",
    profile_image: ''
  };
};