// African countries with their states/regions and major cities
// This enables cascading dropdowns for location selection

export interface LocationData {
  states: {
    [state: string]: string[] // state -> cities
  }
}

export type AfricanCountry = keyof typeof AFRICAN_LOCATIONS

export const AFRICAN_LOCATIONS: Record<string, LocationData> = {
  'Nigeria': {
    states: {
      'Abia': ['Aba', 'Umuahia', 'Ohafia'],
      'Adamawa': ['Yola', 'Mubi', 'Numan', 'Jimeta'],
      'Akwa Ibom': ['Uyo', 'Eket', 'Ikot Ekpene', 'Oron'],
      'Anambra': ['Awka', 'Onitsha', 'Nnewi', 'Ekwulobia'],
      'Bauchi': ['Bauchi', 'Azare', 'Misau', 'Jama\'are'],
      'Bayelsa': ['Yenagoa', 'Brass', 'Ogbia'],
      'Benue': ['Makurdi', 'Gboko', 'Otukpo', 'Katsina-Ala'],
      'Borno': ['Maiduguri', 'Biu', 'Damboa', 'Gwoza'],
      'Cross River': ['Calabar', 'Ogoja', 'Ikom', 'Obudu'],
      'Delta': ['Asaba', 'Warri', 'Sapele', 'Ughelli', 'Agbor'],
      'Ebonyi': ['Abakaliki', 'Afikpo', 'Onueke'],
      'Edo': ['Benin City', 'Auchi', 'Ekpoma', 'Uromi'],
      'Ekiti': ['Ado-Ekiti', 'Ikere-Ekiti', 'Ijero-Ekiti'],
      'Enugu': ['Enugu', 'Nsukka', 'Agbani', 'Oji River'],
      'FCT': ['Abuja', 'Gwagwalada', 'Kuje', 'Bwari'],
      'Gombe': ['Gombe', 'Kumo', 'Billiri', 'Kaltungo'],
      'Imo': ['Owerri', 'Orlu', 'Okigwe', 'Oguta'],
      'Jigawa': ['Dutse', 'Hadejia', 'Gumel', 'Kazaure'],
      'Kaduna': ['Kaduna', 'Zaria', 'Kafanchan', 'Kagoro'],
      'Kano': ['Kano', 'Wudil', 'Rano', 'Gaya'],
      'Katsina': ['Katsina', 'Daura', 'Funtua', 'Malumfashi'],
      'Kebbi': ['Birnin Kebbi', 'Argungu', 'Yauri', 'Zuru'],
      'Kogi': ['Lokoja', 'Okene', 'Kabba', 'Idah'],
      'Kwara': ['Ilorin', 'Offa', 'Jebba', 'Lafiagi'],
      'Lagos': ['Lagos', 'Ikeja', 'Lekki', 'Victoria Island', 'Ikorodu', 'Badagry', 'Epe'],
      'Nasarawa': ['Lafia', 'Keffi', 'Akwanga', 'Nasarawa'],
      'Niger': ['Minna', 'Bida', 'Kontagora', 'Suleja'],
      'Ogun': ['Abeokuta', 'Sagamu', 'Ijebu-Ode', 'Ota'],
      'Ondo': ['Akure', 'Ondo City', 'Owo', 'Ikare'],
      'Osun': ['Osogbo', 'Ile-Ife', 'Ilesa', 'Ede', 'Ikire'],
      'Oyo': ['Ibadan', 'Ogbomoso', 'Oyo', 'Iseyin', 'Saki'],
      'Plateau': ['Jos', 'Bukuru', 'Pankshin', 'Shendam'],
      'Rivers': ['Port Harcourt', 'Bonny', 'Okrika', 'Degema'],
      'Sokoto': ['Sokoto', 'Tambuwal', 'Wurno', 'Gwadabawa'],
      'Taraba': ['Jalingo', 'Wukari', 'Bali', 'Takum'],
      'Yobe': ['Damaturu', 'Potiskum', 'Gashua', 'Nguru'],
      'Zamfara': ['Gusau', 'Kaura Namoda', 'Talata Mafara', 'Anka'],
    }
  },
  'Ghana': {
    states: {
      'Greater Accra': ['Accra', 'Tema', 'Madina', 'Teshie'],
      'Ashanti': ['Kumasi', 'Obuasi', 'Ejisu', 'Mampong'],
      'Western': ['Sekondi-Takoradi', 'Tarkwa', 'Axim'],
      'Eastern': ['Koforidua', 'Nkawkaw', 'Akosombo'],
      'Central': ['Cape Coast', 'Winneba', 'Kasoa', 'Elmina'],
      'Northern': ['Tamale', 'Yendi', 'Damongo'],
      'Volta': ['Ho', 'Hohoe', 'Keta', 'Aflao'],
      'Brong-Ahafo': ['Sunyani', 'Techiman', 'Berekum'],
      'Upper East': ['Bolgatanga', 'Navrongo', 'Bawku'],
      'Upper West': ['Wa', 'Tumu', 'Lawra'],
    }
  },
  'South Africa': {
    states: {
      'Gauteng': ['Johannesburg', 'Pretoria', 'Soweto', 'Sandton', 'Midrand'],
      'Western Cape': ['Cape Town', 'Stellenbosch', 'Paarl', 'George'],
      'KwaZulu-Natal': ['Durban', 'Pietermaritzburg', 'Richards Bay', 'Newcastle'],
      'Eastern Cape': ['Port Elizabeth', 'East London', 'Mthatha', 'Grahamstown'],
      'Free State': ['Bloemfontein', 'Welkom', 'Kroonstad'],
      'Mpumalanga': ['Nelspruit', 'Witbank', 'Secunda', 'Middelburg'],
      'Limpopo': ['Polokwane', 'Thohoyandou', 'Tzaneen', 'Mokopane'],
      'North West': ['Rustenburg', 'Klerksdorp', 'Potchefstroom', 'Mahikeng'],
      'Northern Cape': ['Kimberley', 'Upington', 'Springbok'],
    }
  },
  'Kenya': {
    states: {
      'Nairobi': ['Nairobi', 'Westlands', 'Karen', 'Eastleigh'],
      'Mombasa': ['Mombasa', 'Nyali', 'Likoni'],
      'Kisumu': ['Kisumu', 'Ahero', 'Maseno'],
      'Nakuru': ['Nakuru', 'Naivasha', 'Gilgil'],
      'Eldoret': ['Eldoret', 'Kitale', 'Webuye'],
      'Kiambu': ['Thika', 'Ruiru', 'Kikuyu', 'Limuru'],
      'Machakos': ['Machakos', 'Athi River', 'Kangundo'],
      'Nyeri': ['Nyeri', 'Karatina', 'Othaya'],
    }
  },
  'Cameroon': {
    states: {
      'Centre': ['Yaoundé', 'Mbalmayo', 'Obala'],
      'Littoral': ['Douala', 'Edéa', 'Nkongsamba'],
      'West': ['Bafoussam', 'Dschang', 'Mbouda'],
      'North West': ['Bamenda', 'Kumbo', 'Ndop'],
      'South West': ['Buea', 'Limbe', 'Kumba'],
      'East': ['Bertoua', 'Batouri', 'Abong-Mbang'],
      'South': ['Ebolowa', 'Kribi', 'Sangmélima'],
      'Adamawa': ['Ngaoundéré', 'Meiganga', 'Tibati'],
      'North': ['Garoua', 'Guider', 'Figuil'],
      'Far North': ['Maroua', 'Kousséri', 'Mokolo'],
    }
  },
  'Egypt': {
    states: {
      'Cairo': ['Cairo', 'Heliopolis', 'Nasr City', 'Maadi'],
      'Giza': ['Giza', '6th of October City', 'Sheikh Zayed'],
      'Alexandria': ['Alexandria', 'Borg El Arab'],
      'Luxor': ['Luxor', 'Karnak'],
      'Aswan': ['Aswan', 'Kom Ombo'],
      'Port Said': ['Port Said'],
      'Suez': ['Suez'],
      'Sharm El Sheikh': ['Sharm El Sheikh', 'Dahab'],
      'Hurghada': ['Hurghada', 'El Gouna'],
    }
  },
  'Morocco': {
    states: {
      'Casablanca-Settat': ['Casablanca', 'Mohammedia', 'El Jadida'],
      'Rabat-Salé-Kénitra': ['Rabat', 'Salé', 'Kénitra', 'Témara'],
      'Marrakech-Safi': ['Marrakech', 'Safi', 'Essaouira'],
      'Fès-Meknès': ['Fès', 'Meknès', 'Ifrane'],
      'Tanger-Tétouan-Al Hoceïma': ['Tangier', 'Tétouan', 'Chefchaouen'],
      'Souss-Massa': ['Agadir', 'Taroudant', 'Tiznit'],
      'Oriental': ['Oujda', 'Nador', 'Berkane'],
    }
  },
  'Senegal': {
    states: {
      'Dakar': ['Dakar', 'Pikine', 'Guédiawaye', 'Rufisque'],
      'Thiès': ['Thiès', 'Mbour', 'Tivaouane'],
      'Saint-Louis': ['Saint-Louis', 'Richard Toll', 'Dagana'],
      'Diourbel': ['Diourbel', 'Touba', 'Mbacké'],
      'Kaolack': ['Kaolack', 'Nioro du Rip', 'Guinguinéo'],
      'Ziguinchor': ['Ziguinchor', 'Bignona', 'Oussouye'],
      'Fatick': ['Fatick', 'Foundiougne', 'Gossas'],
    }
  },
  'Ivory Coast': {
    states: {
      'Abidjan': ['Abidjan', 'Cocody', 'Plateau', 'Yopougon', 'Marcory'],
      'Yamoussoukro': ['Yamoussoukro'],
      'Bouaké': ['Bouaké'],
      'San-Pédro': ['San-Pédro', 'Tabou'],
      'Daloa': ['Daloa', 'Issia'],
      'Korhogo': ['Korhogo', 'Ferkessédougou'],
      'Man': ['Man', 'Danané'],
    }
  },
  'Tanzania': {
    states: {
      'Dar es Salaam': ['Dar es Salaam', 'Kinondoni', 'Temeke', 'Ilala'],
      'Arusha': ['Arusha', 'Moshi'],
      'Mwanza': ['Mwanza', 'Sengerema'],
      'Dodoma': ['Dodoma'],
      'Zanzibar': ['Zanzibar City', 'Stone Town'],
      'Mbeya': ['Mbeya', 'Tukuyu'],
      'Morogoro': ['Morogoro', 'Ifakara'],
    }
  },
  'Uganda': {
    states: {
      'Central': ['Kampala', 'Entebbe', 'Mukono', 'Wakiso'],
      'Eastern': ['Jinja', 'Mbale', 'Soroti', 'Tororo'],
      'Western': ['Mbarara', 'Fort Portal', 'Kasese', 'Kabale'],
      'Northern': ['Gulu', 'Lira', 'Arua', 'Kitgum'],
    }
  },
  'Ethiopia': {
    states: {
      'Addis Ababa': ['Addis Ababa', 'Bole', 'Kirkos'],
      'Oromia': ['Adama', 'Bishoftu', 'Jimma', 'Dire Dawa'],
      'Amhara': ['Bahir Dar', 'Gondar', 'Dessie', 'Debre Markos'],
      'Tigray': ['Mekelle', 'Axum', 'Adwa'],
      'SNNPR': ['Hawassa', 'Arba Minch', 'Wolaita Sodo'],
    }
  },
  'Algeria': {
    states: {
      'Algiers': ['Algiers', 'Bab El Oued', 'Hussein Dey'],
      'Oran': ['Oran', 'Ain Turk', 'Es Senia'],
      'Constantine': ['Constantine', 'El Khroub'],
      'Annaba': ['Annaba', 'El Bouni'],
      'Blida': ['Blida', 'Boufarik'],
      'Sétif': ['Sétif', 'El Eulma'],
    }
  },
  'Angola': {
    states: {
      'Luanda': ['Luanda', 'Viana', 'Cacuaco'],
      'Benguela': ['Benguela', 'Lobito'],
      'Huambo': ['Huambo', 'Caála'],
      'Cabinda': ['Cabinda'],
      'Huíla': ['Lubango', 'Matala'],
      'Malanje': ['Malanje'],
    }
  },
  'Zambia': {
    states: {
      'Lusaka': ['Lusaka', 'Chilanga', 'Kafue'],
      'Copperbelt': ['Kitwe', 'Ndola', 'Chingola', 'Mufulira'],
      'Southern': ['Livingstone', 'Choma', 'Mazabuka'],
      'Eastern': ['Chipata', 'Petauke'],
      'Northern': ['Kasama', 'Mbala'],
    }
  },
  'Zimbabwe': {
    states: {
      'Harare': ['Harare', 'Chitungwiza', 'Epworth'],
      'Bulawayo': ['Bulawayo'],
      'Manicaland': ['Mutare', 'Chipinge', 'Rusape'],
      'Mashonaland West': ['Chinhoyi', 'Karoi', 'Kadoma'],
      'Matabeleland North': ['Victoria Falls', 'Hwange'],
      'Midlands': ['Gweru', 'Kwekwe', 'Zvishavane'],
    }
  },
}

// Get list of countries
export const AFRICAN_COUNTRIES = Object.keys(AFRICAN_LOCATIONS).sort()

// Get states for a country
export function getStatesForCountry(country: string): string[] {
  const locationData = AFRICAN_LOCATIONS[country]
  if (!locationData) return []
  return Object.keys(locationData.states).sort()
}

// Get cities for a country and state
export function getCitiesForState(country: string, state: string): string[] {
  const locationData = AFRICAN_LOCATIONS[country]
  if (!locationData) return []
  const cities = locationData.states[state]
  return cities ? [...cities].sort() : []
}

// Check if a country has location data
export function hasLocationData(country: string): boolean {
  return country in AFRICAN_LOCATIONS
}
