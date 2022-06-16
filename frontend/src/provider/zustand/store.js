import create from "zustand";

const useStore = create((set) => ({
  user: undefined,
  setUser: (user) =>
    set((state) => ({
      ...state,
      user: user,
    })),
  resetUser: () => set((state) => ({
    ...state,
    user: undefined,
  })),
}));

export default useStore;