import React from 'react';
import axios from 'axios';
import { useUserContext } from '@/providers/UserProvider';

interface CartButtonProps {
  onClick: () => void;
  itemId: number;
}

const CartButton: React.FC<CartButtonProps> = ({ onClick, itemId}) => {
  const {user} = useUserContext();
  const handleAddToCart = async () => {
    try {
      const response = await axios.post('http://localhost:8080/api/v1/cartItems', {
        cart_id: user?.cart.id,
        item_id: itemId,
        quantity: 1,
      });

      if (response.status === 201) {
        alert('Item added to cart successfully!');
      }
    } catch (error) {
      console.error('Error adding item to cart:', error);
      alert('Failed to add item to cart');
    }
  };

  return (
    <button
      onClick={() => {
        handleAddToCart();
        onClick(); 
      }}
      className="px-4 py-2 bg-blue-600 text-white font-semibold rounded-lg hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-400"
    >
      Add to Cart
    </button>
  );
};

export default CartButton;
