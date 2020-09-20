package hypercard_whiteboard

var eventDrawingHouse = Drawing{
	Xs:        []float32{876.59595, 878.91736, 881.2388, 881.2388, 881.2388, 881.1495, 881.1495, 881.2388, 881.32806, 881.50665, 881.8638, 882.1316, 882.3102, 882.4888, 882.7566, 883.0245, 883.0245, 883.203, 883.5602, 883.7387, 883.9173, 884.1852, 884.453, 884.453, 884.6316, 884.9887, 885.2566, 885.3459, 885.4351, 885.6137, 885.4351, 885.4351, 885.703, 885.6137, 885.6137, 885.2566, 885.078, 884.8994, 884.8102, 884.8994, 884.7209, 884.6316, 884.7209, 884.453, 883.9173, 883.5602, 883.203, 882.9352, 882.6673, 882.57806, 882.3995, 882.04236, 881.7745, 881.50665, 881.32806, 881.1495, 881.2388, 880.9709, 880.34595, 880.07806, 879.81024, 879.3638, 879.00665, 878.91736, 878.82806, 878.56024, 878.20306, 877.84595, 877.75665, 877.5781, 877.22095, 877.0424, 876.9531, 876.86383, 876.86383, 877.31024, 877.22095, 876.77454, 876.68524, 876.68524, 876.9531, 876.9531, 876.77454, 876.68524, 876.68524, 876.77454, 876.77454, 876.59595, 876.68524, 876.86383, 876.86383, 876.86383, 877.0424, 877.0424, 876.86383, 876.86383, 876.9531, 877.13165, 877.5781, 878.38165, 879.27454, 880.4352, 881.5959, 882.8459, 884.0959, 885.7923, 887.578, 889.45294, 891.5065, 893.64935, 895.8815, 897.935, 900.43494, 903.02423, 905.3456, 907.7563, 910.2563, 912.8455, 915.16693, 917.5776, 920.0776, 922.6668, 925.1668, 927.30963, 929.98816, 932.9345, 935.3452, 937.6666, 939.89874, 942.30945, 944.7201, 947.0415, 949.1843, 951.32715, 953.47, 955.6128, 957.5771, 959.1842, 961.0592, 962.93414, 964.63055, 966.1484, 967.9341, 969.3627, 970.5234, 971.9519, 973.2019, 974.45184, 975.61255, 976.77325, 977.7554, 978.64825, 979.5411, 980.34467, 981.1482, 981.7732, 982.1304, 982.6661, 983.1125, 983.4696, 983.8267, 984.6303, 984.8089, 984.6303, 984.8982, 984.9874, 984.9874, 985.0767, 984.8982, 984.8982, 984.9874, 984.9874, 984.9874, 984.6303, 984.7196, 985.52313, 985.4339, 984.9874, 984.9874, 984.9874, 984.9874, 984.9874, 984.8982, 984.8982, 984.9874, 984.9874, 984.9874, 984.4517, 984.4517, 984.9874, 985.0767, 984.541, 985.8803, 985.8803, 984.9874, 984.9874, 984.9874, 985.166, 984.9874, 984.9874, 985.0767, 984.9874, 984.7196, 985.4339, 984.9874, 984.8982, 985.166, 984.8982, 984.9874, 984.8982, 984.8982, 984.8982, 984.9874, 984.8982, 984.9874, 984.9874, 984.8982, 984.8982, 984.8982, 984.8089, 984.8982, 984.8089, 984.7196, 984.7196, 984.6303, 984.541, 984.541, 984.4517, 984.3624, 984.2732, 984.2732, 984.0946, 983.916, 983.7375, 983.4696, 983.2018, 982.9339, 982.8446, 982.6661, 982.3982, 982.1304, 981.8625, 981.6839, 981.4161, 981.2375, 981.1482, 980.96967, 980.7911, 980.7018, 980.43396, 980.2554, 980.1661, 979.89825, 979.80896, 979.36255, 978.73755, 977.93396, 976.41614, 974.63043, 972.1305, 969.4519, 966.95197, 964.09485, 961.14844, 957.9342, 954.80927, 951.41644, 948.3808, 945.7022, 943.0237, 940.61304, 938.55945, 936.5952, 934.6309, 933.0238, 931.6846, 930.4346, 929.27386, 927.75604, 926.7739, 926.32745, 925.5239, 924.80963, 924.18463, 923.7382, 923.2025, 922.8454, 922.5775, 922.1311, 921.774, 921.4168, 921.0597, 920.7026, 920.4347, 920.1669, 919.8097, 919.3633, 918.7383, 918.47046, 918.9169, 918.649, 918.1133, 917.93475, 917.84546, 918.024, 917.84546, 917.4883, 917.39905, 917.30975, 917.13116, 917.0419, 917.13116, 917.0419, 916.9526, 916.77405, 916.77405, 917.13116, 917.0419, 916.77405, 916.77405, 916.68475, 916.86334, 916.77405, 916.59546, 916.50616, 916.23834, 915.88116, 915.52405, 915.07764, 914.54193, 914.18475, 913.5598, 912.8455, 912.04193, 911.2384, 910.5241, 909.8098, 908.917, 907.9349, 906.9527, 906.0599, 904.63135, 902.93494, 901.68494, 900.43494, 899.3635, 898.2028, 896.8636, 895.6136, 894.4529, 893.2029, 891.77435, 890.61365, 889.45294, 888.29224, 887.0423, 885.8816, 884.9887, 884.0959, 883.2923, 882.8459, 882.3995, 882.04236, 881.95306, 881.95306, 881.8638, 882.04236, 882.04236, 881.95306, 881.95306, 882.04236, 882.04236, 882.1316, 882.2209, 882.04236, 882.1316, 882.3102, 882.3102, 881.95306, 881.95306, 882.1316, 882.1316, 882.3102, 882.1316, 882.04236, 881.8638, 882.04236, 882.04236, 882.1316, 882.04236, 882.3102, 882.3995, 882.1316, 882.2209, 882.1316, 882.04236, 882.04236, 882.1316, 882.1316, 882.3102, 882.3995, 882.57806, 882.4888, 882.1316, 882.1316, 882.3995, 882.04236, 882.1316, 882.04236, 882.04236, 882.1316, 882.04236, 882.1316, 882.1316, 882.2209, 882.1316, 882.1316, 882.1316, 882.2209, 882.1316, 882.1316, 882.2209, 882.3102, 882.2209, 882.3102, 882.3102, 882.3102, 882.3995, 882.6673, 883.203, 884.0066, 884.8994, 885.7923, 886.7744, 887.9351, 889.0958, 890.3458, 891.86365, 893.7386, 895.9707, 898.6493, 901.4171, 904.54205, 908.292, 912.57764, 917.0419, 921.6847, 926.86316, 932.2203, 937.57733, 943.0237, 948.7379, 954.45215, 959.8985, 965.25555, 970.3448, 974.98755, 979.36255, 983.6482, 987.5767, 990.61237, 993.2909, 995.9694, 998.3801, 1000.70154, 1002.75507, 1004.4515, 1005.9693, 1007.5764, 1009.4514, 1010.88, 1011.8621, 1012.93353, 1013.6478, 1014.27277, 1014.71924, 1015.16565, 1015.4335, 1015.61206, 1015.8799, 1015.9692, 1015.9692, 1016.0585, 1016.23706, 1016.4156, 1016.5942, 1016.6835, 1016.77277, 1017.0406, 1017.2192, 1017.3085, 1017.3085, 1017.39777, 1017.2192, 1017.2192, 1017.48706, 1017.48706, 1017.39777, 1017.48706, 1017.48706, 1017.48706, 1017.3085, 1017.3085, 1017.2192, 1017.3085, 1017.1299, 1016.5049, 1015.79065, 1014.62994, 1012.93353, 1010.7014, 1008.2014, 1005.25507, 1001.86224, 998.1123, 994.0052, 989.18384, 983.6482, 978.20184, 971.8626, 964.36273, 957.22, 950.25574, 943.113, 936.86304, 931.77386, 925.97034, 920.0776, 914.89905, 909.8991, 904.8992, 900.2564, 895.7922, 891.86365, 888.5601, 885.5244, 883.1138, 881.2388, 879.81024, 878.73883, 878.11383, 877.93524, 877.75665, 877.75665, 877.93524, 878.02454, 878.20306, 878.47095, 878.82806, 879.3638, 879.72095, 879.8995, 880.16736, 880.34595, 880.4352, 880.4352, 880.5245, 880.5245, 880.5245, 880.5245, 880.5245, 880.5245, 880.5245, 880.70306, 880.6138, 880.4352, 880.4352, 880.4352, 880.70306, 880.79236, 880.6138, 880.6138, 880.5245, 880.6138, 880.79236, 880.9709, 881.2388, 881.50665, 881.8638, 882.2209, 882.57806, 883.0245, 883.3816, 883.828, 884.5423, 885.1673, 885.9708, 887.2208, 888.20294, 889.0958, 890.43506, 891.86365, 893.56006, 895.6136, 897.8457, 900.2564, 903.20276, 906.32776, 909.6313, 913.11334, 916.86334, 920.8811, 924.8989, 929.0953, 933.2917, 937.39874, 941.41656, 945.4344, 949.3629, 953.38074, 957.0414, 960.4342, 964.00555, 967.4877, 970.6126, 973.4697, 975.9697, 978.2911, 980.34467, 982.3089, 984.4517, 986.416, 988.2017, 989.80884, 991.23737, 992.48737, 993.5588, 994.3623, 994.9873, 995.4337, 995.7016, 995.8802, 995.9694, 995.8802, 995.7016, 995.6123, 995.523, 995.3445, 995.2552, 994.9873, 994.898, 994.8088, 994.9873, 994.7195, 994.8088, 994.898, 994.898, 994.898, 995.3445, 995.2552, 994.8088, 994.8088, 994.898, 994.8088, 994.8088, 994.898, 994.898, 994.898, 994.898, 994.7195, 994.7195, 994.898, 994.898, 994.8088, 994.8088, 994.7195, 995.0766, 994.6302, 994.8088, 994.7195, 994.6302, 994.5409, 994.5409, 994.5409, 994.5409, 994.3623, 994.4516, 994.5409, 994.5409, 994.3623, 994.3623, 994.4516, 994.4516, 994.3623, 994.2731, 994.2731, 994.5409, 994.9873, 995.6123, 996.148, 996.9516, 997.9337, 998.82654, 999.7194, 1000.79083, 1001.95154, 1003.02295, 1004.18365, 1005.34436, 1006.2372, 1007.3086, 1008.38, 1009.4514, 1010.5228, 1011.50494, 1012.57635, 1013.73706, 1014.8085, 1015.70135, 1016.5942, 1017.5763, 1018.2013, 1018.3799, 1018.8263, 1019.0942, 1019.2727, 1019.2727, 1019.0049, 1018.8263, 1018.4692, 1017.9335, 1017.3085},
	Ys:        []float32{1743.0752, 1743.968, 1745.1287, 1745.5751, 1745.8429, 1745.8429, 1745.9323, 1745.9323, 1746.0215, 1746.0215, 1746.1107, 1746.1107, 1746.0215, 1746.0215, 1746.0215, 1746.0215, 1746.0215, 1745.9323, 1745.8429, 1745.5751, 1745.1287, 1744.3251, 1742.9858, 1740.6646, 1737.8967, 1734.6825, 1730.8434, 1726.7363, 1722.2722, 1717.6294, 1712.8081, 1708.0762, 1703.5227, 1699.0586, 1694.773, 1690.8445, 1687.0054, 1683.1661, 1679.5948, 1676.2914, 1672.8986, 1669.4165, 1665.756, 1662.0953, 1658.3455, 1654.774, 1651.5599, 1648.435, 1645.4886, 1643.1672, 1640.9352, 1638.6139, 1636.9175, 1635.3997, 1634.0604, 1632.8998, 1631.8284, 1630.6676, 1629.507, 1628.3463, 1627.4535, 1626.5607, 1625.5785, 1624.5964, 1623.7036, 1622.8108, 1621.8286, 1621.025, 1620.4001, 1619.6858, 1619.0608, 1618.7037, 1618.5251, 1618.2573, 1617.9895, 1617.8109, 1617.7216, 1617.6323, 1617.6323, 1617.6323, 1617.7216, 1617.7216, 1617.6323, 1617.7216, 1617.7216, 1617.7216, 1617.7216, 1617.8109, 1617.8109, 1617.8109, 1617.8109, 1617.8109, 1617.8109, 1617.8109, 1617.9001, 1617.9001, 1617.9001, 1617.9895, 1618.3466, 1618.7037, 1619.0608, 1619.418, 1619.8644, 1620.2216, 1620.5786, 1621.025, 1621.65, 1622.3643, 1623.0786, 1623.7036, 1624.5071, 1625.3999, 1626.0249, 1626.7391, 1627.3641, 1627.8999, 1628.5249, 1629.0605, 1629.507, 1629.8641, 1630.2212, 1630.3998, 1630.5784, 1630.6676, 1630.5784, 1630.5784, 1630.5784, 1630.3998, 1629.9534, 1629.507, 1628.882, 1628.0785, 1627.0963, 1626.1143, 1624.8643, 1623.7036, 1622.5428, 1621.293, 1620.1322, 1619.0608, 1617.9001, 1616.6501, 1615.4003, 1614.1503, 1612.9003, 1611.6504, 1610.311, 1608.704, 1607.0969, 1605.4897, 1603.8827, 1602.4542, 1601.1149, 1599.865, 1598.615, 1597.4543, 1596.5615, 1595.8472, 1595.2223, 1594.5973, 1593.8829, 1593.258, 1592.9009, 1592.4545, 1592.008, 1591.6509, 1591.4723, 1591.4723, 1591.383, 1591.4723, 1591.4723, 1591.4723, 1591.4723, 1591.4723, 1591.4723, 1591.5616, 1591.6509, 1591.6509, 1591.7402, 1591.8295, 1591.7402, 1591.7402, 1591.7402, 1591.7402, 1591.6509, 1591.7402, 1591.7402, 1591.7402, 1591.7402, 1591.7402, 1591.7402, 1591.7402, 1591.7402, 1591.7402, 1591.8295, 1591.8295, 1591.8295, 1591.8295, 1591.8295, 1591.8295, 1591.8295, 1591.8295, 1591.8295, 1591.8295, 1591.9187, 1591.9187, 1591.9187, 1591.9187, 1592.008, 1592.008, 1591.9187, 1592.008, 1591.9187, 1591.9187, 1591.9187, 1592.008, 1592.0973, 1592.008, 1592.0973, 1592.0973, 1592.0973, 1592.0973, 1592.0973, 1592.1866, 1592.1866, 1592.1866, 1592.2759, 1592.3652, 1592.3652, 1592.4545, 1592.5437, 1592.633, 1592.7223, 1592.9901, 1593.3473, 1593.5259, 1593.7937, 1593.9723, 1594.1509, 1594.4187, 1594.5973, 1594.8651, 1595.0437, 1595.1329, 1595.2223, 1595.3115, 1595.4008, 1595.4901, 1595.4901, 1595.5793, 1595.5793, 1595.5793, 1595.4008, 1595.3115, 1595.0437, 1594.5079, 1593.7045, 1592.2759, 1590.2224, 1587.7224, 1584.8654, 1581.8297, 1578.9727, 1576.1156, 1573.08, 1570.0444, 1566.9194, 1563.8838, 1561.2053, 1558.6161, 1556.2054, 1553.8842, 1551.2949, 1548.7057, 1546.2058, 1543.7952, 1541.3845, 1539.0631, 1536.831, 1534.599, 1532.4562, 1530.3134, 1528.2599, 1526.2957, 1524.51, 1522.9028, 1521.4744, 1520.3137, 1519.153, 1517.9031, 1516.7423, 1515.6709, 1514.8674, 1514.1531, 1513.171, 1511.8318, 1510.2246, 1508.7961, 1507.6355, 1506.7427, 1506.0283, 1505.4033, 1505.0463, 1504.7784, 1504.5105, 1504.2427, 1503.9749, 1503.707, 1503.4391, 1503.2606, 1503.082, 1503.082, 1503.1713, 1503.1713, 1503.2606, 1503.3499, 1503.4391, 1503.4391, 1503.5284, 1503.6177, 1503.707, 1503.7963, 1503.8856, 1503.9749, 1504.4213, 1505.3141, 1506.5641, 1508.0819, 1510.0461, 1512.3674, 1514.5995, 1516.9209, 1519.5101, 1522.3672, 1525.4921, 1528.8849, 1532.1884, 1535.7596, 1539.5095, 1543.438, 1547.6343, 1551.7413, 1555.5806, 1558.9733, 1562.5446, 1566.3838, 1570.4015, 1574.2407, 1578.1692, 1582.3655, 1586.7404, 1591.1152, 1595.4008, 1599.24, 1602.9899, 1606.3827, 1609.4182, 1611.9182, 1613.9717, 1615.6681, 1617.2751, 1618.5251, 1619.418, 1619.7751, 1619.9536, 1619.9536, 1619.8644, 1619.8644, 1619.8644, 1619.7751, 1619.7751, 1619.6858, 1619.6858, 1619.5966, 1619.5966, 1619.5966, 1619.5966, 1619.5966, 1619.5966, 1619.5072, 1619.418, 1619.3287, 1619.2394, 1619.1501, 1619.1501, 1619.1501, 1619.2394, 1619.1501, 1619.1501, 1619.1501, 1619.1501, 1619.2394, 1619.2394, 1619.2394, 1619.2394, 1619.2394, 1619.1501, 1619.1501, 1619.2394, 1619.2394, 1619.1501, 1619.1501, 1619.1501, 1619.1501, 1619.2394, 1619.2394, 1619.3287, 1619.418, 1619.3287, 1619.3287, 1619.3287, 1619.3287, 1619.3287, 1619.2394, 1619.2394, 1619.2394, 1619.3287, 1619.2394, 1619.2394, 1619.2394, 1619.3287, 1619.3287, 1619.3287, 1619.3287, 1619.418, 1619.5072, 1619.5966, 1619.6858, 1619.8644, 1620.2216, 1620.5786, 1621.025, 1621.4714, 1622.0072, 1622.6322, 1623.3464, 1624.3285, 1625.4893, 1626.6499, 1627.9891, 1629.5962, 1631.6498, 1634.0604, 1636.7389, 1640.0424, 1643.6136, 1647.185, 1650.9349, 1654.774, 1658.9705, 1663.0774, 1666.9166, 1670.4879, 1673.8807, 1677.2734, 1680.5769, 1683.7911, 1686.7374, 1689.5945, 1692.3623, 1694.5944, 1696.648, 1698.9692, 1701.3799, 1703.7905, 1706.1119, 1708.1654, 1710.0404, 1711.9153, 1713.8795, 1715.8438, 1717.5402, 1718.8794, 1719.8616, 1720.4865, 1720.8436, 1721.1115, 1721.4686, 1721.7365, 1722.0043, 1722.0936, 1722.3615, 1722.54, 1722.8079, 1723.0757, 1723.3436, 1723.7007, 1723.9685, 1724.0579, 1724.2365, 1724.3257, 1724.4149, 1724.4149, 1724.4149, 1724.4149, 1724.5043, 1724.5043, 1724.4149, 1724.4149, 1724.4149, 1724.4149, 1724.4149, 1724.5043, 1724.5043, 1724.5043, 1724.5935, 1724.7721, 1725.1293, 1725.8435, 1726.5577, 1727.1827, 1727.8077, 1728.4327, 1729.0577, 1729.772, 1730.3076, 1730.9326, 1731.8254, 1732.629, 1733.7004, 1735.0397, 1736.3789, 1737.8075, 1739.5038, 1741.2894, 1742.8073, 1743.968, 1744.6823, 1745.3965, 1746.2001, 1746.9143, 1747.4501, 1748.075, 1748.7893, 1749.4143, 1750.0392, 1750.575, 1751.0214, 1751.3785, 1751.6464, 1751.7356, 1751.825, 1751.7356, 1751.6464, 1751.557, 1751.557, 1751.4678, 1751.2, 1750.932, 1750.6642, 1750.3964, 1750.1285, 1750.0392, 1749.95, 1749.95, 1749.95, 1749.95, 1749.95, 1749.95, 1749.95, 1749.95, 1749.95, 1749.95, 1749.95, 1749.95, 1749.95, 1749.95, 1749.95, 1749.95, 1749.95, 1749.95, 1749.8607, 1749.8607, 1749.8607, 1749.6821, 1749.4143, 1749.2357, 1748.9679, 1748.7, 1748.4321, 1748.075, 1747.7179, 1747.4501, 1747.0929, 1746.5573, 1746.0215, 1745.5751, 1744.8608, 1744.0573, 1743.2537, 1742.093, 1740.5752, 1738.7896, 1736.7361, 1734.5039, 1732.004, 1729.3256, 1726.3792, 1723.0757, 1719.5936, 1715.8438, 1711.826, 1707.5404, 1702.987, 1698.255, 1693.7015, 1689.416, 1685.0411, 1680.4877, 1676.0234, 1671.9165, 1667.8987, 1663.881, 1660.0419, 1656.2026, 1652.4528, 1648.7921, 1645.31, 1641.9174, 1638.5245, 1635.2211, 1632.007, 1628.7034, 1625.6678, 1622.8108, 1620.1322, 1618.0787, 1616.2037, 1614.5074, 1613.1681, 1612.0968, 1611.2932, 1610.6682, 1610.1326, 1609.8646, 1609.5968, 1609.4182, 1609.329, 1609.329, 1609.2397, 1609.0612, 1608.9718, 1608.9718, 1608.9718, 1608.9718, 1608.9718, 1608.8826, 1608.8826, 1608.9718, 1608.9718, 1609.0612, 1609.1504, 1609.0612, 1608.9718, 1609.0612, 1609.0612, 1608.9718, 1608.9718, 1609.0612, 1609.1504, 1609.1504, 1609.0612, 1608.9718, 1608.9718, 1608.9718, 1608.9718, 1608.9718, 1608.9718, 1608.9718, 1608.8826, 1608.7933, 1608.704, 1608.4362, 1608.2576, 1608.1683, 1608.1683, 1608.079, 1608.1683, 1608.2576, 1608.2576, 1608.3468, 1608.4362, 1608.5254, 1608.704, 1608.8826, 1609.1504, 1610.0432, 1611.8289, 1614.3289, 1617.6323, 1621.4714, 1625.7571, 1630.3105, 1634.864, 1639.4174, 1644.328, 1649.3279, 1654.3276, 1658.9705, 1663.3453, 1667.4523, 1671.2915, 1674.9521, 1678.6127, 1682.2733, 1685.8446, 1689.0588, 1692.1837, 1695.0408, 1697.8978, 1700.487, 1702.987, 1705.3083, 1707.6298, 1709.594, 1711.201, 1712.6296, 1713.8795, 1714.9509, 1715.7545, 1716.5581, 1717.0938},
	Pressures: []int32{1510, 1445, 1365, 1368, 1433, 1484, 1508, 1530, 1567, 1633, 1725, 1826, 1948, 2094, 2243, 2373, 2479, 2588, 2709, 2825, 2936, 3055, 3164, 3239, 3306, 3380, 3442, 3482, 3510, 3542, 3569, 3601, 3631, 3654, 3671, 3687, 3694, 3690, 3698, 3710, 3720, 3732, 3743, 3755, 3769, 3776, 3779, 3780, 3793, 3808, 3821, 3828, 3836, 3843, 3850, 3857, 3874, 3878, 3874, 3880, 3887, 3886, 3889, 3903, 3911, 3911, 3914, 3918, 3912, 3913, 3924, 3931, 3929, 3929, 3933, 3917, 3921, 3951, 3959, 3963, 3971, 3977, 3979, 3982, 3985, 3990, 3991, 3986, 3985, 3984, 3981, 3976, 3962, 3953, 3955, 3941, 3915, 3903, 3900, 3901, 3907, 3914, 3910, 3911, 3913, 3913, 3911, 3911, 3912, 3911, 3912, 3918, 3916, 3912, 3917, 3921, 3921, 3921, 3925, 3927, 3927, 3925, 3926, 3932, 3933, 3933, 3934, 3932, 3935, 3936, 3936, 3939, 3944, 3951, 3951, 3949, 3952, 3962, 3973, 3979, 3978, 3981, 3966, 3966, 3983, 3985, 3987, 3988, 3995, 4000, 4005, 3997, 3999, 4002, 4008, 4013, 4004, 4011, 4010, 4014, 4015, 4026, 4041, 4042, 4047, 4049, 4047, 4047, 4051, 4055, 4063, 4066, 4070, 4079, 4086, 4072, 4074, 4091, 4092, 4093, 4093, 4092, 4094, 4093, 4090, 4089, 4087, 4083, 4082, 4086, 4087, 4095, 4077, 4078, 4091, 4092, 4094, 4095, 4094, 4095, 4095, 4095, 4095, 4095, 4095, 4095, 4095, 4095, 4095, 4095, 4095, 4095, 4095, 4095, 4095, 4095, 4095, 4095, 4095, 4095, 4095, 4095, 4095, 4095, 4095, 4095, 4095, 4095, 4095, 4095, 4095, 4102, 4102, 4108, 4108, 4115, 4123, 4123, 4123, 4123, 4123, 4123, 4123, 4123, 4123, 4123, 4123, 4123, 4123, 4123, 4123, 4123, 4132, 4132, 4142, 4142, 4142, 4142, 4151, 4158, 4158, 4158, 4158, 4166, 4175, 4175, 4175, 4175, 4175, 4175, 4175, 4175, 4175, 4175, 4185, 4185, 4185, 4185, 4185, 4196, 4206, 4215, 4223, 4230, 4236, 4236, 4236, 4236, 4236, 4236, 4236, 4236, 4236, 4236, 4236, 4236, 4236, 4236, 4236, 4236, 4243, 4243, 4249, 4249, 4249, 4249, 4249, 4249, 4249, 4249, 4249, 4249, 4249, 4249, 4256, 4264, 4271, 4271, 4271, 4277, 4277, 4277, 4277, 4277, 4277, 4277, 4277, 4282, 4282, 4282, 4282, 4282, 4282, 4282, 4282, 4282, 4288, 4288, 4296, 4305, 4305, 4305, 4313, 4320, 4326, 4326, 4326, 4326, 4326, 4326, 4326, 4326, 4333, 4341, 4100, 4101, 4095, 4095, 4102, 4108, 4108, 4113, 4113, 4113, 4113, 4113, 4113, 4119, 4119, 4119, 4119, 4119, 4126, 4134, 4134, 4134, 4134, 4134, 4134, 4134, 4134, 4134, 4134, 4134, 4134, 4134, 4134, 4134, 4134, 4134, 4134, 4134, 4134, 4134, 4134, 4134, 4134, 4134, 4134, 4134, 4134, 4134, 4134, 4134, 4134, 4134, 4134, 4134, 4134, 4143, 4143, 4143, 4143, 4143, 4143, 4143, 4143, 4143, 4143, 4143, 4143, 4143, 4151, 4151, 4151, 4151, 4151, 4151, 4151, 4151, 4151, 4151, 4151, 4151, 4151, 4151, 4151, 4151, 4151, 4151, 4160, 4160, 4168, 4175, 4175, 4175, 4175, 4183, 4183, 4183, 4183, 4183, 4190, 4190, 4190, 4190, 4190, 4190, 4196, 4201, 4201, 4201, 4201, 4207, 4214, 4222, 4222, 4231, 4241, 4250, 4250, 4258, 4258, 4258, 4258, 4258, 4258, 4258, 4258, 4258, 4258, 4258, 4258, 4258, 4258, 4258, 4258, 4265, 4271, 4271, 4271, 4271, 4271, 4271, 4278, 4278, 4284, 4284, 4284, 4291, 4300, 4300, 4300, 4300, 4300, 4300, 4300, 4300, 4308, 4315, 4321, 4321, 4326, 4326, 4326, 4326, 4326, 4095, 4035, 3998, 3995, 3959, 3942, 3964, 3969, 3961, 3967, 3968, 3970, 3972, 3953, 3941, 3944, 3947, 3949, 3953, 3950, 3946, 3946, 3945, 3938, 3922, 3891, 3846, 3810, 3778, 3757, 3757, 3759, 3770, 3775, 3774, 3772, 3768, 3771, 3773, 3777, 3782, 3788, 3794, 3800, 3806, 3810, 3792, 3795, 3821, 3824, 3825, 3823, 3821, 3821, 3822, 3824, 3825, 3828, 3831, 3835, 3842, 3846, 3849, 3843, 3841, 3840, 3839, 3844, 3842, 3843, 3843, 3840, 3827, 3824, 3830, 3831, 3834, 3831, 3831, 3852, 3864, 3865, 3888, 3879, 3868, 3864, 3849, 3856, 3865, 3873, 3884, 3896, 3906, 3919, 3943, 3956, 3952, 3954, 3962, 3963, 3968, 3977, 3984, 3994, 4004, 4013, 4014, 4016, 4003, 4002, 4020, 4022, 4024, 4025, 4030, 4023, 4028, 4048, 4060, 4077, 4091, 4095, 4095, 4095, 4095, 4095, 4095, 4095, 4095, 4094, 4085, 4088, 4082, 4066, 4065, 4068, 4068, 4069, 4073, 4075, 4072, 4078, 4082, 4080, 4095, 4095, 4091, 4095, 4095, 4095, 4095, 4095, 4100, 4100, 4100, 4100, 4100, 4100, 4104, 4104, 4104, 4104, 4104, 4094, 4092, 4084, 4075, 4063, 4053, 4049, 4046, 4043, 4046, 4039, 4034, 4031, 4016, 4002, 3993, 3979, 3970, 3956, 3939, 3924, 3915, 3901, 3885, 3867, 3853, 3830, 3805, 3780, 3742, 3687, 3633, 3565, 3485, 3397, 3270, 3095, 2847, 2467, 1814, 1297, 44},
	Widths:    []uint32{2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2},
	Color:     Drawing_BLACK,
}
