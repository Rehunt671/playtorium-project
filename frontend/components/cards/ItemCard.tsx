import React from 'react';
import CartButton from '../buttons/CartButton';
import { Item } from '@/types/item';
import { useMutationAddCartItem, useMutationUpdateCartItem, useQueryCartItems } from '@/query/cart';
import { useUserContext } from '@/providers/UserProvider';
import { CartItem } from '@/types/cart_item';

interface ItemCardProps {
  item: Item;
}

const ItemCard: React.FC<ItemCardProps> = ({ item }) => {
  const { id, name, price, image_name, description } = item;
  const { user } = useUserContext();
  const { data: cartItems } = useQueryCartItems(user?.cart.id!);
  const mutationCreateCartItem = useMutationAddCartItem(user?.cart.id!);
  const mutationUpdateCartItem = useMutationUpdateCartItem(user?.cart.id!, id);

  const handleAddToCart = async (itemId: number) => {
    const existingCartItem = cartItems?.find(cartItem => cartItem.item_id === itemId);

    if (existingCartItem) {
      const cartItem: CartItem = {
        id: existingCartItem.id,
        cart_id: user?.cart.id!,
        item_id: itemId,
        quantity: existingCartItem.quantity + 1,
        created_at: existingCartItem.created_at,
        updated_at: new Date().toISOString(),
        item: existingCartItem.item,
      };
      mutationUpdateCartItem.mutateAsync(cartItem);
    } else {
      mutationCreateCartItem.mutateAsync(
        {
          id : 0,
          cart_id: user?.cart.id!,
          item_id: itemId,
          quantity: 1,
        }
      );
    }
  };
  return (
    <div className="bg-white rounded-lg shadow-lg overflow-hidden group hover:shadow-2xl transition-all duration-300 ease-in-out transform hover:scale-105">
      <div className="relative">
      <div
          className="w-full h-56 bg-gray-200 rounded-t-lg mb-4"
          style={{
            background: `url(${image_name ? `/images/${image_name}` : 'https://via.placeholder.com/300x200'}) no-repeat center center / cover`,
          }}
        >
        </div>

        <div className="p-6 pb-16">
          <h2 className="text-xl font-semibold text-gray-800 mb-2 truncate">{name}</h2>
          <p className="text-lg font-medium text-gray-600 mb-4">${price.toFixed(2)}</p>
          <p className="text-sm text-gray-500 mb-4">{description}</p>
          <div className="absolute bottom-6 left-1/2 transform -translate-x-1/2">
            <CartButton itemId={id} onClick={() => handleAddToCart(id)} />
          </div>
        </div>
      </div>
    </div>
  );
};

export default ItemCard;