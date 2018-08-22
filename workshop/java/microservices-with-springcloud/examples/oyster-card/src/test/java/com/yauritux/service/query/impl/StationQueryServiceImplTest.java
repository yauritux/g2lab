package com.yauritux.service.query.impl;

import static org.assertj.core.api.Assertions.assertThat;
import static org.junit.Assert.assertTrue;

import java.util.Map;
import java.util.Optional;

import org.junit.Before;
import org.junit.Test;
import org.junit.runner.RunWith;
import org.mockito.Mockito;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.boot.test.mock.mockito.MockBean;
import org.springframework.test.context.junit4.SpringRunner;

import com.yauritux.DataTestConfiguration;
import com.yauritux.model.entity.Station;
import com.yauritux.repository.StationRepository;
import com.yauritux.service.query.StationQueryService;

/**
 * 
 * @author yauritux
 * @version 1.0.0
 * @since 1.0.0
 *
 */
@RunWith(SpringRunner.class)
@SpringBootTest(classes = DataTestConfiguration.class)
public class StationQueryServiceImplTest {
	
	private static final String[] STATION_NAMES =  {
			"Holborn", "Earl's Court", "Wimbledon", "Hammersmith"
	};

	@MockBean
	private StationRepository stationRepository;
	
	@Autowired
	private Map<String, Station> stations;
	
	private StationQueryService stationQueryService;
	
	@Before
	public void setup() {
		stationQueryService = new StationQueryServiceImpl(stationRepository);
	}
	
	@Test
	public void findByName_existingStation_shouldReturnStation() {
		Mockito.when(stationRepository.findByName(STATION_NAMES[0])).thenReturn(stations.get(STATION_NAMES[0]));
		
		Optional<Station> found = stationQueryService.findByName(STATION_NAMES[0]);
		assertTrue(found.isPresent());
		assertThat(found.get().getName()).isEqualTo(stations.get(STATION_NAMES[0]).getName());
	}
}
