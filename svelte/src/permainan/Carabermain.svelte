<script>
    import { TabContent, TabPane } from 'sveltestrap';
    import Panel from '../components/Panel.svelte'
</script>
<Panel>
    <slot:template slot="header">
        <h1 style="padding:0px;margin:0px;color:white;font-size:15px;">CARA BERMAIN</h1>
    </slot:template>
    <slot:template slot="body">
        <TabContent style="padding: 0px;margin:0px;">
            <TabPane tabId="bermain_432d" tab="4D/3D/2D" active style="padding:20px 5px 20px 22px;font-size:12px;color:#8a8a8a;">
                Menebak 4 angka, 3 angka dan 2 angka
				<br />
				Struktur ABCD<br /><br />

				Misalnya keluar : 4321<br />
				Berarti pemenang untuk<br />
				4D = 4321<br />
				3D = 321<br />
				2D = 21<br />
                2DD = 43<br />
                2DT = 32<br /><br />

				Aturan permainan:<br />
				1. Jika anda membeli diluar dari nomor yang dikeluarkan, berarti anda kalah dan uang tidak dikembalikan sama sekali.<br />
				2. Jika anda membeli masing 100rb untuk angka :<br />
				4D = 4321<br />
				3D = 321<br />
				2D = 21<br /><br />	
				Berarti kemenangan anda adalah :<br />
				4D = 100rb x [Indeks kemenangan untuk 4D]<br />
				3D = 100rb x [Indeks kemenangan untuk 3D]<br />
				2D = 100rb x [Indeks kemenangan untuk 2D]<br /><br />
				(Catatan: nilai bet 100rb tidak dikembalikan ke member)<br />
				(Khusus untuk 4D,3D dan 2D diberikan diskon tambahan)<br />
            </TabPane>
            <TabPane tabId="bermain_colokbebas" tab="COLOK" style="padding:20px 5px 20px 22px;font-size:12px;color:#8a8a8a;">
                COLOK BEBAS
				<br />
				Menebak salah satu angka dari 4D. Posisi angka bisa dimana saja
				<br />
				Struktur ABCD<br /><br />
				
				Analisis I :<br />
				keluar : 4321<br />
				Misalnya pembelian Angka 3 dengan nilai taruhan 100rb.<br />
				Berarti menang : 100rb + [Indeks kemenangan untuk colok angka]<br /><br />
				
				Analisis I:<br />
				keluar : 4331<br />
				Misalnya pembelian Angka 3 dengan nilai taruhan 100rb.<br />
				Berarti menang: 100rb + ([Indeks kemenangan untuk colok Angka] x 2)<br /><br />
				
				Dan seterusnya untuk setiap kembaran yang berhasil ditebak, otomatis mendapat kelipatan [Indeks kemenangan untuk colok angka]
				
				<hr />
				COLOK MACAU
				<br />
				Cara kerja seperti colok angka tapi mesti yang keluar 2 angka dari 4D.
				<br />
				Struktur ABCD<br /><br />
				
				Analisis I :<br />
				keluar : 4321<br />
				Misalnya dibeli 4 dan 2 dengan nilai 100rb.<br />
				Berarti menang: 100rb + [Indeks kemenangan untuk Macau, kategori: 2 digit]<br /><br />
				
				Analisis II :<br />
				keluar : 4321<br />
				Misalnya dibeli 4 dan 6 dengan nilai 100rb.<br />
				Berarti KALAH dan nilai betting tidak dikembalikan)<br /><br />
				
				Analisis III :<br />
				keluar : 4331<br />
				Misalnya dibeli 4 dan 3 dengan nilai 100rb.<br />
				Berarti menang: 100rb + [Indeks kemenangan untuk Macau, kategori: 3 digit]<br /><br />
				
				Analisis IV :<br />
				keluar : 4334<br />
				Misalnya dibeli 4 dan 3 dengan nilai 100rb.<br />
				Berarti menang: 100rb + [Indeks kemenangan untuk Macau, kategori: 3 digit]
				
				<hr />
				COLOK NAGA : 3 angka dari 4D.
				<br />
				Struktur ABCD<br /><br />
				
				Analisis I :<br />
				keluar : 4321<br />
				Misalnya dibeli 4,2 dan 3 dengan nilai 100rb.<br />
				Berarti menang: karena keluar 3 digit,angka 4,2 dan 3.<br />
				100rb + [Indeks kemenangan untuk colok naga , kategori: 3 digit]<br /><br />
				
				Analisis II :<br />
				keluar : 4321<br />
				Misalnya dibeli 4,2 dan 6 dengan nilai 100rb.<br />
				karena keluar hanya 2 digit angka 4 dan 2, dan angka 6 tidak muncul berarti kalah<br /><br />
				
				Analisis III :<br />
				keluar : 4331<br />
				Misalnya dibeli 4,3 dan 3 dengan nilai 100rb.<br />
				Berarti menang: karena keluar 3 digit,angka 4,3 dan 3.<br />
				100rb + [Indeks kemenangan untuk colok naga , kategori: 3 digit]<br /><br />
				
				Analisis IV :<br />
				keluar : 4334<br />
				Misalnya dibeli 4,3 dan 3 dengan nilai 100rb.<br />
				Berarti menang: 100rb + [Indeks kemenangan untuk colok naga, kategori: 4 digit]
				
				<hr />
				COLOK JITU, Menebak satu angka pada posisi tertentu dari 4D.
				<br />
				Struktur ABCD<br /><br />
				
				Analisis I :<br />
				keluar : 4321<br />
				Misalnya dibeli 4 pada posisi AS dengan nilai 100rb.<br />
				Berarti menang: 100rb + [Indeks kemenangan untuk colok jitu]<br /><br />
				
				Analisis II :<br />
				keluar : 4331<br />
				Misalnya dibeli 3 pada posisi KOP dengan nilai 100rb.<br />
				Berarti menang: 100rb + [Indeks kemenangan untuk colok jitu]. Hasilnya sama dengan analisis I karena hanya memperhatikan posisi yang dipasang.<br /><br />
				
				Analisis III :<br />
				keluar : 4331<br />
				Misalnya dibeli 4 pada posisi EKOR dengan nilai 100rb.<br />
				Berarti kalah. Biarpun nilai 4 keluar pada posisi AS tapi tidak akan mepengaruhi pemilihan di pososi EKOR
            </TabPane>
            <TabPane tabId="bermain_5050umum" tab="50 - 50" style="padding:20px 5px 20px 22px;font-size:12px;color:#8a8a8a;">
                50-50 UMUM, Permainan ganjil/genap, besar/kecil, dan tengah/ tepi UMUM
				<br />
				Struktur: CD ( hanya dua angka belakang yang dihitung)<br /><br />
				
				Menebak ganjil/genap dan besar/kecil dari posisi:<br />
				C=KEPALA<br />
				D=EKOR<br />
				Besar/Kecil: 0-4=kecil, 5-9=besar<br />
				Ganjil/Genap : 1=ganjil, 2=genap dan seterusnya<br /><br />
				Tengah/Tepi : Tengah: angka 25 s/d 74. TEPI : angka 75 s/d 99, dan 00s/d 24<br /><br />
				
				Analisis I :<br />
				Keluar : 6789, Berarti hasil 89<br />
				Berarti pemenang adalah yang memilih :<br />
				BESAR<br />
				GANJIL<br />
				TEPI
				
				<hr />
				50-50 SPECIAL, Menebak ganjil/genap, besar/kecil
				<br />
				Permainan ini sangat menarik karena pasarannya naik turun sesuai keinginan market pada waktu tersebut. Dengan demikian, nilai pembelian dipengaruhi kei (pasaran)
				<br /><br />
				Struktur ABCD<br /><br />
				
				Menebak ganjil/genap dan besar/kecil dari posisi:<br />
				A=AS<br />
				B=KOP<br />
				C=KEPALA<br />
				D=EKOR<br />
				
				Besar/Kecil : 0-4=kecil, 5-9=besar<br />
				Ganjil/Genap : 1=ganjil, 2=genap dan seterusnya<br /><br />
				
				Analisis I :<br />
				Keluar : 4327<br />
				Berarti pemenang adalah yang memilih :<br />
				AS GENAP/KECIL<br />
				KOP GANJIL/KECIL<br />
				KEPALA GENAP/KECIL<br />
				EKOR GANJIL/BESAR<br />
				Misal anda membeli dengan dana Rp.100rb untuk AS Genap, menang = 100rb + [indeks pasaran AS Genap 50-50]<br />
				Atau :<br />
				Jika membeli dengan dana Rp.100rb untuk Ekor Ganjil, menang = 100rb + [indeks pasaran Ekor Ganjil 50-50]<br />
				Atau :<br />
				Jika membeli dengan dana Rp.100rb untuk AS Ganjil, kalah = 100rb + [indeks pasaran AS Ganjil 50-50]
				
				<hr />
				50-50 KOMBINASI, MONO STEREO
				<br />
				Menebak dari posisi Depan,Tengah,Belakang.
				<br /><br />
				Contoh no 1234<br />
				Yang dimaksud dengan Posisi Depan adalah 2 no terdepan yaitu 12<br />
				Yang dimaksud dengan Posisi Tengah adalah 2 no ditengah yaitu 23<br />
				Yang dimaksud dengan Posisi Belakang adalah 2 no terbelakang yaitu 34<br />
				STEREO = Terdapat Ganjil dan Genap<br />
				MONO = Terdapat 1 pasang Ganjil atau 1 pasang Genap<br /><br />
				
				Analisis I : Beli Posisi Depan<br />
				Keluar : 1234<br />
				Yang menjadi pedoman adalah posisi Depan, berarti 12<br />
				12 => 1=ganjil dan 2=genap , berarti hasil = STEREO<br /><br />
				
				Analisis II : Beli Posisi Tengah<br />
				Keluar : 4326<br />
				Yang menjadi pedoman adalah posisi Tengah, berarti 32<br />
				32 => 3=ganjil dan 2=genap , berarti hasil = STEREO<br /><br />
				
				Analisis III : Beli Posisi Belakang<br />
				Keluar : 4533<br />
				Yang menjadi pedoman adalah posisi Belakang, berarti 33<br />
				33 => 3=ganjil dan 3=ganjil , berarti hasil = MONO<br />
				Jika dilakukan pembelian dengan 100rb dan menang maka: Menang = 100rb + [Indeks kemenangan untuk MONO STEREO]<br /><br />
				
				KEMBANG KEMPIS KEMBAR<br />
				Struktur=ABCD<br />
				Jika Menebang posisi Depan maka yang menjadi fokus adalah AB<br />
				Jika Menebang posisi Tengah maka yang menjadi fokus adalah BC<br />
				Jika Menebang posisi Belakang maka yang menjadi fokus adalah CD<br />
				KEMBANG jika A &#8249; B ataupun B &#8249; C ataupun C &#8249; D<br />
				KEMPIS jika A <span style='font-size:12px;'>&#8250</span> B ataupun B <span style='font-size:12px;'>&#8250</span> C ataupun C <span style='font-size:12px;'>&#8250</span> D<br />
				KEMBAR jika A = B ataupun B = C ataupun C = D<br /><br />
				
				Analisis I : Beli Posisi Depan<br />
				Keluar : 4321<br />
				Yang menjadi pedoman adalah posisi Depan, berarti 43<br />
				43 => 4 <span style='font-size:12px;'>&#8250</span>; 3, hasil = KEMPIS<br /><br />

				Analisis II : Beli Posisi Tengah<br />
				Keluar : 4236<br />
				Yang menjadi pedoman adalah posisi Tengah, berarti 23<br />
				23 => 2 &#8249; 3, hasil = KEMBANG<br /><br />

				Analisis III : Beli Posisi Belakang:<br />
				Keluar : 4099<br />
				Yang menjadi pedoman adalah posisi Belakang, berarti 99<br />
				99 => Hasil = KEMBAR<br /><br />

				Jika dilakukan pembelian dengan 100rb dan menang maka:<br />
				Menang = 100rb + [Indeks kemenangan untuk KEMBANG/KEMPIS/KEMBAR]
            </TabPane>
            <TabPane tabId="bermain_macaukombinasi" tab="KOMBINASI" style="padding:20px 5px 20px 22px;font-size:12px;color:#8a8a8a;">
                Struktur ABCD<br />
				AB = DEPAN, BC = TENGAH, CD = BELAKANG<br />
				Besar/Kecil: 0-4=kecil, 5-9=besar<br />
				Ganjil/Genap : 1=ganjil, 2=genap dan seterusnya<br /><br />

				Anda dapat menebak Genap/Ganjil, Besar/Kecil<br />
				dari 2 kombinasi antara DEPAN, TENGAH, BELAKANG<br /><br />

				Analisis : keluar nomor 1845<br /><br />

				berarti pemenang untuk :<br />
				DEPAN Kecil/Genap<br />
				TENGAH Besar/Genap<br />
				BELAKANG Kecil/Ganjil<br /><br />

				Misalnya anda membeli BELAKANG KECIl dan GANJIL seharga 100rb,<br />
				maka menang = 100rb + [indeks kemenangan untuk kombinasi 2]<br />
				atau :<br />
				jika membeli DEPAN KECIL dan GENAP seharga 100rb,<br />
				maka menang = 100rb + [indeks kemenangan untuk kombinasi 2]<br />
				atau :<br />
				jika membeli TENGAH KECIL dan GENAP seharga 100rb, berarti KALAH<br />
				( Anda harus menebak keduanya dengan Benar diantara DEPAN,TENGAH,BELAKANG agar Menang )<br />
            </TabPane>
            <TabPane tabId="bermain_dasar" tab="DASAR" style="padding:20px 5px 20px 22px;font-size:12px;color:#8a8a8a;">
                Menebak ganjil/genap dan besar/kecil dari penjumlah angka-angka 2D.<br />
				Nilai pembelian ditentukan pasaran (kei) pada saat itu.<br /><br />

				Struktur CD (2 angka terakhir)<br />
				Kecil = angka 0-4<br />
				Besar = angka 5-9<br />
				Ganjil = 1,3,5,7,9<br />
				Genap = 0,2,4,6,8<br /><br />

				Analisis I :<br />
				Keluar : 1234,<br />
				3+4 = 7<br />
				Berarti keluar : Ganjil dan Besar<br /><br />

				Analisis II :<br />
				Keluar : 5678,<br />
				7+8 = 15<br />
				Karena angka 15 lebih besar dari 9, kembali dihitung 1+5=6<br />
				Berarti keluar : Genap dan Besar<br /><br />

				Analisis III :<br />
				Keluar : 1204,<br />
				0+4 = 4<br />
				Berarti keluar : Genap dan Kecil<br />
				Misal anda membeli dengan Rp.100rb untuk Genap, menang = 100rb + [indeks menang untuk Dasar]<br /><br />

				NB: Indeks menang dan diskon dapat dilihat di bagian Peraturan
            </TabPane>
            <TabPane tabId="bermain_shio" tab="SHIO" style="padding:20px 5px 20px 22px;font-size:12px;color:#8a8a8a;">
                Struktur ABCD<br />
				Menebak SHIO dari posisi 2D, SHIO merupakan 12 lambang kelahiran dalam penanggalan China. Dalam permainan ini, setiap lambang diwakili dengan satu nomor.<br /><br />

				Analisis I :<br />
				Keluar : 4321<br />
				Permainan ini hanya memperhatikan posisi 2D, berarti yang dipedomanin = 21<br />
				Hasil = 21-12 = 9 (shio disesuaikan dengan tabel diatas)<br />
				catatan: nilai yang dikurangi merupakan kelipatan 12.<br /><br />

				Jika dilakukan pembelian dengan 100rb dan menang maka:<br />
				Menang = 100rb + [Indeks kemenangan untuk SHIO]<br /><br />

				NB: Indeks menang dan diskon dapat dilihat di bagian Peraturan
            </TabPane>
        </TabContent>
    </slot:template>
</Panel>