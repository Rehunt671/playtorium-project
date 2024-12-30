import React from 'react';
import { CartItem } from '@/types/cart_item';
import { useMutationUpdateCartItem, useMutationDeleteCartItem } from '@/query/cart';
import { useUserContext } from '@/providers/UserProvider';

interface CartItemCardProps {
  cartItem: CartItem;
}

const CartItemCard: React.FC<CartItemCardProps> = ({ cartItem }) => {
  const { user } = useUserContext();
  const { item } = cartItem;
  const mutationUpdateCartItem = useMutationUpdateCartItem(user?.cart.id!, cartItem.id);
  const mutationDeleteCartItem = useMutationDeleteCartItem(user?.cart.id!, cartItem.id);

  const handleIncreaseQuantity = () => {
    mutationUpdateCartItem.mutateAsync({ ...cartItem, quantity: cartItem.quantity + 1 });
  };

  const handleDecreaseQuantity = () => {
    if (cartItem.quantity > 1) {
      mutationUpdateCartItem.mutateAsync({ ...cartItem, quantity: cartItem.quantity - 1 });
    } else {
      handleRemoveItem();
    }
  };

  const handleRemoveItem = () => {
    mutationDeleteCartItem.mutateAsync();
  };

  return (
    <div className="bg-white rounded-lg shadow-lg p-6 mb-6 flex border border-gray-200 hover:shadow-xl transition-shadow duration-300">
      <div className="w-28 h-28 mr-6">
        <img
          src={`/images/${item?.image_name}`} 
          alt={item?.name}
          className="w-full h-full object-cover rounded-lg"
        />
      </div>
      {/* Product Info */}
      <div className="flex flex-col justify-between w-full">
        <h2 className="text-xl font-semibold text-gray-800 mb-2">{item?.name}</h2>
          {/* Product Description */}
        <p className="text-sm text-gray-500 mb-4">{item?.description}</p>
        <p className="text-lg font-medium text-gray-600 mb-2">${item?.price.toFixed(2)}</p>
        <div className="flex items-center justify-between space-x-4 mt-2">
          {/* Quantity Controls */}
          <div className="flex items-center space-x-4">
            <button
              onClick={handleDecreaseQuantity}
              className="px-4 py-2 bg-gray-200 hover:bg-gray-300 rounded-lg transition-colors"
            >
              -
            </button>
            <span className="text-lg font-semibold text-gray-800">{cartItem.quantity}</span>
            <button
              onClick={handleIncreaseQuantity}
              className="px-4 py-2 bg-gray-200 hover:bg-gray-300 rounded-lg transition-colors"
            >
              +
            </button>
          </div>

          {/* Remove Button */}
          <button
            onClick={handleRemoveItem}
            className="px-4 py-2 bg-red-500 hover:bg-red-600 text-white rounded-lg transition-colors"
          >
            Remove
          </button>
        </div>
      </div>
    </div>
  );
};

export default CartItemCard;
