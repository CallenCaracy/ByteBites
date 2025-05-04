import { useQuery, useMutation } from "@apollo/client";
import { GET_CART, GET_CART_ITEMS, CREATE_ORDER_FROM_CART } from "../graphql/Orderqueries";
import Navbar from "../components/NavBar";
import "../styles/main.css";
import { useNavigate, useParams } from "react-router-dom";

import React, { useState } from "react";

const YourCart: React.FC = () => {
    const navigate = useNavigate();
    const { userId } = useParams<{ userId: string }>();

    const [orderType, setOrderType] = useState("dine-in");
    const [deliveryAddress, setDeliveryAddress] = useState("");
    const [specialRequests, setSpecialRequests] = useState("");

    const {
        data: cartData,
        loading: cartLoading,
        error: cartError,
    } = useQuery(GET_CART, {
        variables: { user_id: userId },
    });

    const cart = cartData?.getCart;

    const {
        data: cartItemsData,
        loading: cartItemsLoading,
        error: cartItemsError,
        called: cartItemsCalled,
    } = useQuery(GET_CART_ITEMS, {
        variables: { cart_id: cart?.id },
        skip: !cart?.id,
    });

    const [createOrderFromCart, { loading: creatingOrder, error: createOrderError, data: orderData }] =
        useMutation(CREATE_ORDER_FROM_CART);

    const cartItems = cartItemsData?.getCartItemsByCartId;

    if (cartLoading || (cart?.id && !cartItemsCalled) || cartItemsLoading) {
        return <p>Loading your cart...</p>;
    }

    if (cartError) {
        return <p>Error loading cart: {cartError.message}</p>;
    }
    if (cartItemsError) {
        return <p>Error loading cart items: {cartItemsError.message}</p>;
    }

    if (!cartItems || cartItems.length === 0) {
        return (
            <>
                <Navbar />
                <div className="p-8 text-center">
                    <h2 className="text-2xl font-bold">Your Cart</h2>
                    <p className="text-gray-500 mt-4">Your cart is empty. Start ordering now!</p>
                    <p className="text-gray-500 mt-4">If not showing. Try reloading!</p>
                </div>
            </>
        )
    }

    const handleCreateOrder = async () => {
        try {
            if (creatingOrder) return;
            const { data } = await createOrderFromCart({
                variables: {
                    cartID: cart.id,
                    userID: userId,
                    orderType,
                    deliveryAddress: orderType === "delivery" ? deliveryAddress : null,
                    specialRequests,
                },
            });
    
            console.log("Order created:", data.createOrderFromCart);
            navigate(`/payment/${userId}`);
            setTimeout(() => window.location.reload(), 100);
        } catch (err) {
            console.error("Failed to create order:", err);
        }
    };

    return (
        <>
            <Navbar />
            <div className="p-8 max-w-4xl mx-auto">
                <h2 className="text-3xl font-bold mb-6 text-black">Your Cart</h2>
                <div className="max-h-[500px] overflow-y-auto border rounded-lg p-4 shadow-md bg-white scrollbar-hide">
                    {cartItems?.map((item: any) => (
                        <div
                            key={item.id}
                            className="border-b last:border-none pb-3 mb-3 flex justify-between items-center"
                        >
                            <div className="space-y-1">
                                <p className="text-lg font-semibold">Menu Item ID: {item.menu_item_id}</p>
                                <p className="text-sm text-gray-600">Quantity: {item.quantity}</p>
                                <p className="text-sm text-gray-600">
                                    Customizations: {item.customizations || "None"}
                                </p>
                                <p className="text-sm text-gray-600">
                                    Added on: {new Date(item.created_at).toLocaleString()}
                                </p>
                            </div>
                            <div className="text-right">
                                <p className="text-lg font-bold text-green-700">₱{item.price.toFixed(2)}</p>
                            </div>
                        </div>
                    ))}
                </div>

                <div className="mt-6 border-t pt-6">
                    <h3 className="text-xl font-semibold mb-4">Place Your Order</h3>
                    <form
                        className="fixed bottom-0 left-0 w-full bg-white shadow-lg p-4 flex flex-wrap gap-4 items-end border-t z-50"
                        onSubmit={(e) => {
                            e.preventDefault();
                            createOrderFromCart({
                                variables: {
                                    cartID: cart.id,
                                    userID: userId,
                                    orderType,
                                    deliveryAddress,
                                    specialRequests,
                                },
                            });
                        }}
                    >
                        <div className="flex-1 min-w-[200px]">
                            <label className="block text-sm font-semibold text-gray-600">Order Type</label>
                            <select
                                title="Order Type"
                                value={orderType}
                                onChange={(e) => setOrderType(e.target.value)}
                                className="w-full p-2 border rounded-lg text-black"
                            >
                                <option value="dine-in">Dine-in</option>
                                <option value="takeout">Takeout</option>
                                <option value="delivery">Delivery</option>
                            </select>
                        </div>

                        <div className="flex-1 min-w-[200px]">
                            <label className="block text-sm font-semibold text-gray-600">Delivery Address</label>
                            <input
                                type="text"
                                value={deliveryAddress}
                                onChange={(e) => setDeliveryAddress(e.target.value)}
                                className="w-full p-2 border rounded-lg text-black"
                                placeholder="e.g. 123 Cyber Lane"
                            />
                        </div>

                        <div className="flex-1 min-w-[200px]">
                            <label className="block text-sm font-semibold text-gray-600">Special Requests</label>
                            <input
                                type="text"
                                value={specialRequests}
                                onChange={(e) => setSpecialRequests(e.target.value)}
                                className="w-full p-2 border rounded-lg text-black"
                                placeholder="e.g., Allergies, etc."
                            />
                        </div>

                        <div className="flex items-end gap-2">
                            <button
                                type="submit"
                                onClick={handleCreateOrder}
                                className="bg-green-600 text-white px-6 py-2 rounded-lg hover:bg-green-700 transition"
                                disabled={creatingOrder}
                            >
                                {creatingOrder ? "Placing..." : "Place Order"}
                            </button>
                            <button
                                type="button"
                                onClick={() => window.location.reload()}
                                className="bg-gray-400 text-white px-6 py-2 rounded-lg hover:bg-gray-500 transition"
                            >
                                Cancel
                            </button>
                        </div>
                    </form>

                    {createOrderError && (
                        <p className="text-red-600 mt-4">Error: {createOrderError.message}</p>
                    )}
                    {orderData && (
                        <p className="text-green-700 font-medium mt-4">
                            Order placed! Order ID: {orderData.createOrderFromCart.id}
                        </p>
                    )}
                </div>
            </div>
        </>
    );
};

export default YourCart;
