import React, { createContext, useReducer, useContext } from 'react';

const CartStateCtx = createContext([]);
const CartDispCtx  = createContext(() => {});

function cartReducer(state, action) {
    switch (action.type) {
        case 'ADD_ITEM':
            const exists = state.find(i => i.id === action.payload.id);
            if (exists) {
                return state.map(i =>
                    i.id === action.payload.id
                        ? { ...i, qty: i.qty + 1 }
                        : i
                );
            }
            return [...state, { ...action.payload, qty: 1 }];

        case 'CLEAR_CART':
            return [];

        default:
            return state;
    }
}

export function CartProvider({ children }) {
    const [state, dispatch] = useReducer(cartReducer, []);
    return (
        <CartDispCtx.Provider value={dispatch}>
            <CartStateCtx.Provider value={state}>
                {children}
            </CartStateCtx.Provider>
        </CartDispCtx.Provider>
    );
}

export const useCart = () => useContext(CartStateCtx);
export const useCartDispatch = () => useContext(CartDispCtx);
