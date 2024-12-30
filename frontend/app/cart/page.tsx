'use client';
import React from 'react';
import { useUserContext } from '@/providers/UserProvider';
import CartItemCard from '@/components/cards/CartItemCard';
import { useQueryCartDetail } from '@/query/cart';
import { useQueryDiscounts } from '@/query/discount';
import DiscountCard from '@/components/cards/DiscountCard';

const CartPage: React.FC = () => {
  const { user } = useUserContext();
  const { data: cartDetail } = useQueryCartDetail(user?.id!);
  const { data: discounts } = useQueryDiscounts();

  if (!cartDetail || !cartDetail.cart) {
    return (
      <div className="min-h-screen flex items-center justify-center bg-gray-100">
        <div className="text-lg text-gray-600">Loading...</div>
      </div>
    );
  }

  const { cart, total_price, user_points_used } = cartDetail;

  return (
    <div className="min-h-screen bg-gray-50 py-8 px-4 sm:px-6 lg:px-8">
      <div className="max-w-7xl mx-auto">
        <h1 className="text-3xl font-bold text-center text-gray-800 mb-8">Your Cart</h1>

        {/* Render Cart Items */}
        <div className="space-y-4">
          {cart.items.map(cartItem => (
            <CartItemCard key={cartItem.id} cartItem={cartItem} />
          ))}
        </div>

        {/* Render Discounts */}
        {cart.items.length > 0 &&
          discounts?.map((discount) => (
            <DiscountCard
              key={discount.id}
              discount={discount}
            />
          ))}

        {/* Cart Summary */}
        <div className="mt-8 p-4 bg-white rounded-lg shadow-md">
          <div className="flex justify-between items-center text-lg font-semibold">
            <span>Total:</span>
            <span>${total_price.toFixed(2)}</span>
          </div>
          <div className="flex justify-between items-center text-lg font-semibold mt-2">
            <span>Points Used:</span>
            <span>{user_points_used}</span>
          </div>
        </div>
      </div>
    </div>
  );
};

export default CartPage;
