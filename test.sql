Begin
    If Input_Text ~ 'Æon Heart'
		Or Input_Text ~ 'Diamond Cut Heart'
		Or Input_Text ~ 'Facet'
		Or Input_Text ~ 'Signature Bow' Then
        Return 'Golden Way of Life';

	If Input_Text ~'Signature'
		Or Input_Text ~ 'Fingerprint Disc Ring'
		Or Input_Text ~ 'One Zodiac Necklace'
		Or Input_Text ~ 'One Diamond Necklace'
		Or Input_Text ~ 'Two Diamond Necklace'
		Or Input_Text ~ 'Three Diamond Necklace'
		Or Input_Text ~ 'Blend Necklace'
		Or Input_Text ~ 'Diamond Chain'
		Or Input_Text ~ 'One Blend Oval Satin Bracelet'
		Or Input_Text ~ 'One Blend Chain Bracelet'
		Or Input_Text ~ 'Tricolore'
		Or Input_Text ~'One Blend Earring'
		Or Input_Text ~'Diamond Earring'
		Or Input_Text ~'One Blend Chain Earring'
		Or Input_Text ~'One Diamond Chain Earring'
		Or Input_Text ~'Birthstone'
		Or Input_Text ~'From Satin To Chain' Then
        Return 'Signature';
    
	Elsif Input_Text ~'Cords'
		Or Input_Text ~'Kordeln'
		Or Input_Text ~'Koordjes'
		Or Input_Text ~'Nieuwe Satijnkoordjes'
		Or Input_Text ~'Five Extra Cords' Then
   		Return 'Cords';
	
	Elsif Input_Text ~'Giftcard' Then
        Return 'Giftcard';
	
	Elsif Input_Text ~'Forevermore' Then
        Return 'Forevermore';

	Elsif Input_Text ~'Western' Then
        Return 'Western';
	
	Elsif Input_Text ~'Obi' Then
        Return 'Obi';

	Elsif Input_Text ~'Æon'
		Or Input_Text ~'Aeon'
		Or Input_Text ~'AEON'
		Or Input_Text ~'ÆON'
			Or Input_Text ~'√ÜON' Then
        Return 'Aeon';
	
	Elsif Input_Text ~'Benji' Then
        Return 'Benji';
	
	Elsif Input_Text ~'Circle Of Love' Then
        Return 'Circle Of Love';
	
	Elsif Input_Text ~'Connecting Dots' Then
        Return 'Connecting Dots';
	
	Elsif Input_Text ~'Entangle' Then
        Return 'Entangle';
	
	Elsif Input_Text ~'Exagoni' 
		Or Input_Text ~'Tilt' Then
        Return 'Exagoni';		
    
	Elsif Input_Text ~'Flag' Then
        Return 'Flag';
		
	Elsif Input_Text ~'Zodiac Earring'
		Or Input_Text ~'Moony Earring'
		Or Input_Text ~'Yin Yang Earring'
		Or Input_Text ~'Bolt Earring'
		Or Input_Text ~'Comet Earring'
		Or Input_Text ~'Saturn Earring' Then
        Return 'Galaxy';

	Elsif Input_Text ~'Iniemini Hoop'
		Or Input_Text ~'Mini Hoop'
		Or Input_Text ~'Bold Hoop'
		Or Input_Text ~'Chunky Hoop'
		Or Input_Text ~'Croissant Hoop' 
		Or Input_Text ~'Small Hoop' Then
        Return 'Hoops';
		
	Elsif Input_Text ~'Jubilee' Then
        Return 'Jubilee';
	
	Elsif Input_Text ~'Key To My Heart' Then
        Return 'Key To My Heart';

Elsif Input_Text ~'Love Chain' 
		Or Input_Text ~'LOVE CHAIN' Then
        Return 'Love Chain';

	
	Elsif Input_Text ~'Lucky Leaf' Then
        Return 'Lucky Leaf';
	
	--Elsif Input_Text ~'Fingerprint'
	--	Or Input_Text ~'Ash Jewellery' Then
    	--Return 'Memorial';	
	
	Elsif Input_Text ~'Memory Lane'
		Or Input_Text ~'Sparkle Snake'
		Or Input_Text ~'Lucky Lux'
		Or Input_Text ~'Talismom'
		Or Input_Text ~'Disc'
		Or Input_Text ~'Horizontal Edge'
		Or Input_Text ~'Open Tag'
		Or Input_Text ~'Lucky Lux' Then
        Return 'Memorylane';	
	
	Elsif Input_Text ~'Oblique' Then
        Return 'Oblique';	
	
	Elsif Input_Text ~'Baby Love' 
		Or Input_Text ~'Baby Poetry' 
		Or Input_Text ~'Starfish' Then
        Return 'Off Spring';

	Elsif Input_Text ~'Pavé Initial' 
		Or Input_Text ~'Pav√©'
		Or Input_Text ~'Pavé'
		Or Input_Text ~'Pavé Round'
		Or Input_Text ~'Pavé Oval'
		Or Input_Text ~'Fully Pavé' 
		Or Input_Text ~'Pavé Symbol' 
		Or Input_Text ~'Pavé Cross Necklace' Then
        Return 'Pave';

	Elsif Input_Text ~'Queen' Then
        Return 'Queen Of Hearts';

	Elsif Input_Text ~'Scripted' Then
        Return 'Scripted';
	
	Elsif Input_Text ~'Signet' Then
        Return 'Signet';
	
	Elsif Input_Text ~'Spheres' Then
        Return 'Spheres';
	
	Elsif Input_Text ~'Hedge' Then
        Return 'Hedge';
		
	Elsif Input_Text ~'Cross Charm'
		Or Input_Text ~'Olive'
		Or Input_Text ~'Arrow'
		Or Input_Text ~'Pfeil'
		Or Input_Text ~'Feather'
		Or Input_Text ~'Angelwing'
		Or Input_Text ~'Heart Charm'
		Or Input_Text ~'Lucky Charm'
		Or Input_Text ~'Kind Charm'
		Or Input_Text ~'Zodiac Charm'
		Or Input_Text ~'Extra Love Charm'
		Or Input_Text ~'Vertical'
		Or Input_Text ~'Charmed'
		Or Input_Text ~'Chain Of Love'
		Or Input_Text ~'Anchor'
		Or Input_Text ~'Bicolor'
		Or Input_Text ~'Mesh'
		Or Input_Text ~'Poetry'
		Or Input_Text ~'Verse' Then
        Return 'Tokens Of Love';
	
	Elsif Input_Text ~'Unchained' Then
        Return 'Unchained';
	
	Elsif Input_Text ~'Unity' Then
        Return 'Unity';
	
	Elsif Input_Text ~'Mayfair'
		Or Input_Text ~'Victoria'
		Or Input_Text ~'Chelsea' Then
        Return 'Your Heart My Home';
	
	Elsif Input_Text ~'Giftcard'
		Or Input_Text ~'Schiet'
		Or Input_Text ~'Workshop'
		Or Input_Text ~'Reparation'
		Or Input_Text ~'Reparatie'
		Or Input_Text ~'Snowglobe'
		Or Input_Text ~'Ring Sizer'
		Or Input_Text ~'Personalised'
		Or Input_Text ~'custom'
		Or Input_Text ~'hoodie'
		Or Input_Text ~'Surcharge'
		Or Input_Text ~'Engraved'
		Or Input_Text ~'Laser engraving'
		Or Input_Text ~'Laser Engraving'
		Or Input_Text ~'Paw-' Then
   		Return 'No Brand';

	Else
        Return Null;
    End If;
End;