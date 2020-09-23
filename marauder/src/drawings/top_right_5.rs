use crate::proto::whiteboard::{drawing::Color, Drawing};

pub fn f(c: Color) -> Vec<Drawing> {
    vec![
        Drawing {
            xs: vec![
                1230.4308, 1185.2528, 1185.0742, 1184.8063, 1184.6278, 1184.4492, 1184.4492,
                1184.4492, 1184.6278, 1184.8956, 1185.2528, 1185.8777, 1186.6813, 1187.9313,
                1189.4491, 1191.1455, 1193.0205, 1194.8955, 1197.0383, 1199.0919, 1201.7704,
                1204.6276, 1207.4846, 1210.6989, 1213.3774, 1215.9667, 1218.5559, 1220.9666,
                1223.288, 1225.5201, 1227.4844, 1229.3594, 1231.0558, 1232.6628, 1234.0914,
                1235.1628, 1236.0557,
            ],
            ys: vec![
                70.5337, 25.445702, 25.9814, 26.606382, 27.14208, 27.588497, 27.409931, 27.052797,
                26.874231, 26.784948, 26.606382, 26.427816, 25.892117, 25.356419, 24.820719,
                24.285019, 23.838604, 23.392189, 22.85649, 22.32079, 21.874374, 21.338675,
                21.070826, 20.713694, 20.62441, 20.35656, 19.999428, 19.642296, 19.195879,
                18.92803, 18.660181, 18.303047, 17.76735, 17.23165, 16.517385, 16.160252,
                15.713837,
            ],
            pressures: vec![
                0 / 2,
                2853 / 2,
                2937 / 2,
                3099 / 2,
                3330 / 2,
                3503 / 2,
                3578 / 2,
                3598 / 2,
                3604 / 2,
                3627 / 2,
                3678 / 2,
                3732 / 2,
                3774 / 2,
                3806 / 2,
                3858 / 2,
                3934 / 2,
                4000 / 2,
                4047 / 2,
                4088 / 2,
                4095 / 2,
                4098 / 2,
                4057 / 2,
                3977 / 2,
                3986 / 2,
                4000 / 2,
                3996 / 2,
                3985 / 2,
                3980 / 2,
                3970 / 2,
                3953 / 2,
                3916 / 2,
                3845 / 2,
                3684 / 2,
                3393 / 2,
                2906 / 2,
                1946 / 2,
                94 / 2,
            ],
            widths: vec![
                2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2,
                2, 2, 2, 2, 2, 2, 2, 2, 2,
            ],
            color: c as i32,
        },
        Drawing {
            xs: vec![
                1197.2169, 1197.2169, 1197.0383, 1196.8597, 1196.5919, 1196.1455, 1195.7883,
                1195.3419, 1195.1633, 1194.8955, 1194.6277, 1194.3597, 1194.1812, 1194.0027,
                1193.9133, 1193.7347, 1193.4669, 1193.2883, 1193.1097, 1192.9313, 1192.8419,
                1192.7527, 1192.8419, 1192.8419, 1192.6633, 1192.6633, 1192.3955, 1192.2169,
                1192.1277, 1191.8599, 1191.5027, 1191.1455, 1190.8777, 1190.6991, 1190.5205,
                1190.4313, 1190.0741, 1189.9849, 1189.8063, 1189.8063, 1190.0741, 1190.1635,
                1190.2527, 1190.3419, 1190.5205, 1190.7885, 1191.2349, 1191.7705, 1192.3955,
                1192.9313, 1193.5563, 1194.2705, 1194.9847, 1195.6991, 1196.7705, 1198.3776,
                1199.8062, 1201.6812, 1203.4668, 1205.1632, 1207.2168, 1209.2703, 1211.3239,
                1213.1095, 1214.8059, 1216.3237, 1217.6631, 1218.5559, 1219.3595, 1219.9845,
                1220.4309, 1220.8773, 1221.2345, 1221.2345, 1221.3237, 1221.1451, 1220.8773,
                1220.6987, 1220.3416, 1220.0737, 1219.6273, 1219.0023, 1218.1095, 1216.7703,
                1214.8059, 1212.2167, 1208.8239, 1204.9846, 1200.9668, 1196.7705, 1192.7527,
                1189.092, 1185.6099, 1182.5742, 1180.0742, 1178.1992, 1176.86, 1176.235, 1176.4136,
                1176.7708,
            ],
            ys: vec![
                19.46373, 19.999428, 20.535128, 21.338675, 22.05294, 22.499357, 23.035055,
                23.481472, 24.01717, 24.731436, 25.356419, 26.338533, 27.231363, 28.392046,
                29.64201, 30.891973, 32.052654, 33.30262, 34.55258, 36.070396, 37.320362, 38.48104,
                39.463158, 40.355988, 41.24882, 42.05237, 43.034485, 44.37373, 45.80226, 47.40936,
                49.016457, 50.177135, 51.33782, 52.052082, 52.58778, 53.12348, 53.212765,
                53.302048, 53.302048, 53.302048, 53.12348, 53.12348, 53.12348, 52.76635, 52.230648,
                51.4271, 50.26642, 49.10574, 48.39147, 47.76649, 47.40936, 47.052227, 46.516525,
                45.80226, 44.998714, 44.552296, 44.10588, 44.016598, 44.195164, 44.284447,
                44.37373, 44.37373, 44.82015, 45.355846, 45.891544, 46.784378, 47.498642,
                48.212906, 49.10574, 50.087852, 51.06997, 52.230648, 53.569897, 54.909145,
                56.337673, 57.58764, 58.65904, 59.64115, 60.4447, 61.426815, 62.14108, 62.855347,
                63.56961, 64.01603, 64.55173, 65.176704, 65.53384, 66.06954, 66.426674, 66.60524,
                66.783806, 66.783806, 67.05165, 67.05165, 67.05165, 67.05165, 66.60524, 65.53384,
                63.56961, 61.158966,
            ],
            pressures: vec![
                1795 / 2,
                1848 / 2,
                1941 / 2,
                2169 / 2,
                2499 / 2,
                2769 / 2,
                2927 / 2,
                3068 / 2,
                3188 / 2,
                3278 / 2,
                3362 / 2,
                3433 / 2,
                3492 / 2,
                3555 / 2,
                3613 / 2,
                3664 / 2,
                3693 / 2,
                3699 / 2,
                3712 / 2,
                3730 / 2,
                3748 / 2,
                3767 / 2,
                3788 / 2,
                3813 / 2,
                3834 / 2,
                3849 / 2,
                3866 / 2,
                3886 / 2,
                3907 / 2,
                3927 / 2,
                3943 / 2,
                3955 / 2,
                3961 / 2,
                3969 / 2,
                3982 / 2,
                3998 / 2,
                4004 / 2,
                4025 / 2,
                4046 / 2,
                4064 / 2,
                4071 / 2,
                4073 / 2,
                4071 / 2,
                4068 / 2,
                4066 / 2,
                4072 / 2,
                4070 / 2,
                4073 / 2,
                4075 / 2,
                4090 / 2,
                4095 / 2,
                4095 / 2,
                4095 / 2,
                4095 / 2,
                4095 / 2,
                4095 / 2,
                4095 / 2,
                4095 / 2,
                4095 / 2,
                4095 / 2,
                4095 / 2,
                4095 / 2,
                4095 / 2,
                4095 / 2,
                4095 / 2,
                4095 / 2,
                4095 / 2,
                4095 / 2,
                4095 / 2,
                4098 / 2,
                4098 / 2,
                4098 / 2,
                4102 / 2,
                4102 / 2,
                4108 / 2,
                4115 / 2,
                4115 / 2,
                4115 / 2,
                4115 / 2,
                4115 / 2,
                4115 / 2,
                4123 / 2,
                4132 / 2,
                4132 / 2,
                4142 / 2,
                4153 / 2,
                4165 / 2,
                4165 / 2,
                4178 / 2,
                4192 / 2,
                4192 / 2,
                4208 / 2,
                4225 / 2,
                4225 / 2,
                4243 / 2,
                4259 / 2,
                4274 / 2,
                4288 / 2,
                3542 / 2,
                1034 / 2,
            ],
            widths: vec![
                2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2,
                2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2,
                2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2,
                2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2,
            ],
            color: c as i32,
        },
    ]
}
