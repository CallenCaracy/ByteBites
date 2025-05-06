import React from 'react';
import { useQuery, useSubscription } from '@apollo/client';
import { GET_ORDER_QUEUES, ORDER_QUEUE_CREATED } from "../graphql/Orderqueries";
import Navbar from '../components/NavBar';

const OrderQueues: React.FC = () => {
    const { loading, error, data } = useQuery(GET_ORDER_QUEUES);
  
    useSubscription(ORDER_QUEUE_CREATED, {
      onData: ({ client, data }) => {
        const newOrderQueue = data?.data?.orderQueueCreated;
        if (!newOrderQueue) return;
  
        console.log("Received new order queue:", newOrderQueue);
  
        const existing = client.readQuery({ query: GET_ORDER_QUEUES });
        if (existing) {
          console.log("Existing order queues:", existing.orderQueues);
        }
  
        if (existing && !existing.orderQueues?.some((queue: any) => queue.id === newOrderQueue.id)) {
          console.log("New order queue detected, updating cache...");
          client.writeQuery({
            query: GET_ORDER_QUEUES,
            data: {
              orderQueues: [newOrderQueue, ...existing.orderQueues],
            },
          });
        }
      },
      onError: (error) => {
        console.error("Subscription error:", error);
      },
    });
  
    if (loading) return <p className="text-center text-gray-600">Loading...</p>;
    if (error) return <p className="text-center text-red-500">Error loading data.</p>;
  
    return (
        <><Navbar />
        <div className="h-screen flex flex-col">
            <div className="container p-8 flex-1 overflow-y-auto scrollbar-hide">
                <h1 className="text-3xl font-semibold text-gray-800 mb-6">Order Queues</h1>

                {/* Display the list of order queues */}
                <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
                    {data?.orderQueues?.map((orderQueue: any) => (
                        <div
                            key={orderQueue.id}
                            className="group bg-white p-4 rounded-lg shadow-md cursor-pointer hover:bg-blue-950 transition"
                        >
                            <div className="flex justify-between items-center">
                                <h2 className="text-xl font-semibold text-gray-700 group-hover:text-white">
                                    Order ID: {orderQueue.orderId}
                                </h2>
                                <p className="text-xl font-semibold text-gray-700 group-hover:text-white">
                                    {orderQueue.status}
                                </p>
                            </div>
                            <p className="text-gray-600 group-hover:text-gray-300">
                                Created At: {new Date(orderQueue.createdAt).toLocaleString()}
                            </p>
                        </div>
                    ))}
                </div>
            </div>
        </div></>
    );
  };
  
  export default OrderQueues;